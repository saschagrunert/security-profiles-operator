package flect

var pluralRules = []rule{}

// AddPlural adds a rule that will replace the given suffix with the replacement suffix.
func AddPlural(suffix string, repl string) {
	pluralMoot.Lock()
	defer pluralMoot.Unlock()
	pluralRules = append([]rule{{
		suffix: suffix,
		fn: func(s string) string {
			s = s[:len(s)-len(suffix)]
			return s + repl
		},
	}}, pluralRules...)

	pluralRules = append([]rule{{
		suffix: repl,
		fn:     noop,
	}}, pluralRules...)
}

var singleToPlural = map[string]string{
	"aircraft":    "aircraft",
	"alias":       "aliases",
	"alumna":      "alumnae",
	"alumnus":     "alumni",
	"analysis":    "analyses",
	"antenna":     "antennas",
	"antithesis":  "antitheses",
	"apex":        "apexes",
	"appendix":    "appendices",
	"axis":        "axes",
	"bacillus":    "bacilli",
	"bacterium":   "bacteria",
	"basis":       "bases",
	"beau":        "beaus",
	"bison":       "bison",
	"bureau":      "bureaus",
	"bus":         "buses",
	"campus":      "campuses",
	"caucus":      "caucuses",
	"child":       "children",
	"château":     "châteaux",
	"circus":      "circuses",
	"codex":       "codices",
	"concerto":    "concertos",
	"corpus":      "corpora",
	"crisis":      "crises",
	"criterion":   "criteria",
	"curriculum":  "curriculums",
	"datum":       "data",
	"deer":        "deer",
	"diagnosis":   "diagnoses",
	"die":         "dice",
	"dwarf":       "dwarves",
	"ellipsis":    "ellipses",
	"equipment":   "equipment",
	"erratum":     "errata",
	"fez":         "fezzes",
	"fish":        "fish",
	"focus":       "foci",
	"foo":         "foos",
	"foot":        "feet",
	"formula":     "formulas",
	"fungus":      "fungi",
	"genus":       "genera",
	"goose":       "geese",
	"graffito":    "graffiti",
	"grouse":      "grouse",
	"half":        "halves",
	"halo":        "halos",
	"hoof":        "hooves",
	"human":       "humans",
	"hypothesis":  "hypotheses",
	"index":       "indices",
	"information": "information",
	"jeans":       "jeans",
	"larva":       "larvae",
	"libretto":    "librettos",
	"loaf":        "loaves",
	"locus":       "loci",
	"louse":       "lice",
	"matrix":      "matrices",
	"medium":      "media",
	"minutia":     "minutiae",
	"money":       "money",
	"moose":       "moose",
	"mouse":       "mice",
	"nebula":      "nebulae",
	"news":        "news",
	"nucleus":     "nuclei",
	"oasis":       "oases",
	"octopus":     "octopi",
	"offspring":   "offspring",
	"opus":        "opera",
	"ovum":        "ova",
	"ox":          "oxen",
	"parenthesis": "parentheses",
	"person":      "people",
	"phenomenon":  "phenomena",
	"photo":       "photos",
	"phylum":      "phyla",
	"piano":       "pianos",
	"plus":        "pluses",
	"police":      "police",
	"prognosis":   "prognoses",
	"prometheus":  "prometheuses",
	"quiz":        "quizzes",
	"quota":       "quotas",
	"radius":      "radiuses",
	"referendum":  "referendums",
	"ress":        "resses",
	"rice":        "rice",
	"salmon":      "salmon",
	"sex":         "sexes",
	"series":      "series",
	"sheep":       "sheep",
	"shoe":        "shoes",
	"shrimp":      "shrimp",
	"species":     "species",
	"stimulus":    "stimuli",
	"stratum":     "strata",
	"swine":       "swine",
	"syllabus":    "syllabi",
	"symposium":   "symposiums",
	"synapse":     "synapses",
	"synopsis":    "synopses",
	"tableau":     "tableaus",
	"testis":      "testes",
	"thesis":      "theses",
	"thief":       "thieves",
	"tooth":       "teeth",
	"trout":       "trout",
	"tuna":        "tuna",
	"vedalia":     "vedalias",
	"vertebra":    "vertebrae",
	"vertix":      "vertices",
	"vita":        "vitae",
	"vortex":      "vortices",
	"wharf":       "wharves",
	"wife":        "wives",
	"woman":       "women",
	"wolf":        "wolves",
	"you":         "you",
}

var pluralToSingle = map[string]string{}

func init() {
	for k, v := range singleToPlural {
		pluralToSingle[v] = k
	}
}

type singularToPluralSuffix struct {
	singular string
	plural   string
}

var singularToPluralSuffixList = []singularToPluralSuffix{
	{"campus", "campuses"},
	{"person", "people"},
	{"phylum", "phyla"},
	{"randum", "randa"},
	{"actus", "acti"},
	{"adium", "adia"},
	{"basis", "basis"},
	{"child", "children"},
	{"chive", "chives"},
	{"focus", "foci"},
	{"genus", "genera"},
	{"hello", "hellos"},
	{"jeans", "jeans"},
	{"louse", "lice"},
	{"media", "media"},
	{"mouse", "mice"},
	{"movie", "movies"},
	{"oasis", "oasis"},
	{"atum", "ata"},
	{"atus", "atuses"},
	{"base", "bases"},
	{"cess", "cesses"},
	{"dium", "diums"},
	{"eses", "esis"},
	{"half", "halves"},
	{"hive", "hives"},
	{"iano", "ianos"},
	{"irus", "iri"},
	{"isis", "ises"},
	{"leus", "li"},
	{"mnus", "mni"},
	{"move", "moves"},
	{"news", "news"},
	{"odex", "odice"},
	{"oose", "eese"},
	{"ouse", "ouses"},
	{"ovum", "ova"},
	{"shoe", "shoes"},
	{"stis", "stes"},
	{"tive", "tives"},
	{"vice", "vices"},
	{"wife", "wives"},
	{"afe", "aves"},
	{"bfe", "bves"},
	{"box", "boxes"},
	{"cfe", "cves"},
	{"dfe", "dves"},
	{"dge", "dges"},
	{"efe", "eves"},
	{"gfe", "gves"},
	{"hfe", "hves"},
	{"ife", "ives"},
	{"itz", "itzes"},
	{"ium", "ia"},
	{"ize", "izes"},
	{"jfe", "jves"},
	{"kfe", "kves"},
	{"man", "men"},
	{"mfe", "mves"},
	{"nfe", "nves"},
	{"nna", "nnas"},
	{"oaf", "oaves"},
	{"oci", "ocus"},
	{"ode", "odes"},
	{"ofe", "oves"},
	{"pfe", "pves"},
	{"qfe", "qves"},
	{"quy", "quies"},
	{"rfe", "rves"},
	{"sfe", "sves"},
	{"tfe", "tves"},
	{"tum", "ta"},
	{"tus", "tuses"},
	{"ufe", "uves"},
	{"ula", "ulae"},
	{"ula", "ulas"},
	{"uli", "ulus"},
	{"use", "uses"},
	{"uss", "usses"},
	{"vfe", "vves"},
	{"wfe", "wves"},
	{"xfe", "xves"},
	{"yfe", "yves"},
	{"you", "you"},
	{"zfe", "zves"},
	{"by", "bies"},
	{"ch", "ches"},
	{"cy", "cies"},
	{"dy", "dies"},
	{"ex", "ices"},
	{"fy", "fies"},
	{"gy", "gies"},
	{"hy", "hies"},
	{"io", "ios"},
	{"jy", "jies"},
	{"ky", "kies"},
	{"lf", "lves"},
	{"ly", "lies"},
	{"my", "mies"},
	{"ny", "nies"},
	{"py", "pies"},
	{"qy", "qies"},
	{"rf", "rves"},
	{"ry", "ries"},
	{"sh", "shes"},
	{"ss", "sses"},
	{"sy", "sies"},
	{"ty", "ties"},
	{"tz", "tzes"},
	{"va", "vae"},
	{"vy", "vies"},
	{"wy", "wies"},
	{"xy", "xies"},
	{"zy", "zies"},
	{"zz", "zzes"},
	{"o", "oes"},
	{"x", "xes"},
}

func init() {
	for i := len(singularToPluralSuffixList) - 1; i >= 0; i-- {
		AddPlural(singularToPluralSuffixList[i].singular, singularToPluralSuffixList[i].plural)
		AddSingular(singularToPluralSuffixList[i].plural, singularToPluralSuffixList[i].singular)
	}
}
