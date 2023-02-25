/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/bpfrecorder/types"
)

const header = `//go:build linux && !no_bpf
// +build linux,!no_bpf

/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package recorder

`

const (
	buildDir    = "build/"
	bpfObj      = "cli_recorder.bpf.o"
	baseDir     = "internal/pkg/cli/recorder/"
	generatedGo = baseDir + "generated.go"
	btfDir      = "internal/pkg/daemon/bpfrecorder/btf"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	builder := &strings.Builder{}

	if err := generateBpfObj(builder); err != nil {
		return fmt.Errorf("generate bpf object: %w", err)
	}

	if err := generateBtf(builder); err != nil {
		return fmt.Errorf("generate btf: %w", err)
	}

	//nolint:gosec // permissions are fine
	if err := os.WriteFile(
		//nolint:gomnd // filemode is fine
		generatedGo, []byte(builder.String()), 0o644,
	); err != nil {
		return fmt.Errorf("writing generated object: %w", err)
	}
	if err := exec.Command("go", "fmt", generatedGo).Run(); err != nil {
		return fmt.Errorf("format go code: %w", err)
	}
	return nil
}

func generateBpfObj(builder *strings.Builder) error {
	builder.WriteString(header)
	builder.WriteString("var bpfObjects = map[string][]byte{\n")

	for _, arch := range []string{"amd64", "arm64"} {
		builder.WriteString(fmt.Sprintf("%q: {\n", arch))

		file, err := os.ReadFile(filepath.Join(buildDir, bpfObj+"."+arch))
		if err != nil {
			return fmt.Errorf("read bpf object path: %w", err)
		}

		size := len(file)
		for k, v := range file {
			builder.WriteString(fmt.Sprint(v))

			if k < size-1 {
				builder.WriteString(", ")
			}

			if k != 0 && k%16 == 0 {
				builder.WriteString("\n\t")
			}
		}

		builder.WriteString("},\n")
	}

	builder.WriteString("}\n\n")
	return nil
}

func generateBtf(builder *strings.Builder) error {
	builder.WriteString("var btfJSON = `")
	btfs := types.Btf{}

	if err := filepath.Walk(btfDir, func(path string, info fs.FileInfo, retErr error) error {
		if info.IsDir() || filepath.Ext(path) != ".btf" {
			return nil
		}

		// A path should consist of:
		// - the btf dir
		// - the OS
		// - the OS version
		// - the architecture
		// - the btf file containing the kernel version
		pathSplit := strings.Split(path, string(os.PathSeparator))
		const expectedBPFPathLen = 9
		if len(pathSplit) != expectedBPFPathLen {
			return fmt.Errorf("invalid btf path: %s (len = %d)", path, len(pathSplit))
		}

		btfBytes, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read btf file: %w", err)
		}

		os := types.Os(pathSplit[5])
		osVersion := types.OsVersion(pathSplit[6])
		arch := types.Arch(pathSplit[7])
		kernel := types.Kernel(pathSplit[8][0 : len(pathSplit[8])-len(filepath.Ext(pathSplit[8]))])

		if _, ok := btfs[os]; !ok {
			btfs[os] = map[types.OsVersion]map[types.Arch]map[types.Kernel][]byte{}
		}
		if _, ok := btfs[os][osVersion]; !ok {
			btfs[os][osVersion] = map[types.Arch]map[types.Kernel][]byte{}
		}
		if _, ok := btfs[os][osVersion][arch]; !ok {
			btfs[os][osVersion][arch] = map[types.Kernel][]byte{}
		}

		btfs[os][osVersion][arch][kernel] = btfBytes

		return nil
	}); err != nil {
		return fmt.Errorf("walk btf files: %w", err)
	}
	jsonBytes, err := json.MarshalIndent(btfs, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal btf JSON: %w", err)
	}
	builder.Write(jsonBytes)
	builder.WriteString("`\n")

	return nil
}