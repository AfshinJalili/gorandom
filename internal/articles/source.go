package articles

type Source string

const (
	SourceDocs        Source = "docs"
	SourceTour        Source = "tour"
	SourceGoByExample Source = "gobyexample"
	SourcePkg         Source = "pkg"
	SourceBlog        Source = "blog"
)

var Sources = []Source{
	SourceDocs,
	SourceTour,
	SourceGoByExample,
	SourcePkg,
	SourceBlog,
}

func FormatSource(s Source) string {
	switch s {
	case SourceDocs:
		return "Go Docs"
	case SourceTour:
		return "Tour of Go"
	case SourceGoByExample:
		return "Go by Example"
	case SourcePkg:
		return "Standard Library"
	case SourceBlog:
		return "Go Blog"
	default:
		return string(s)
	}
}

func IsValidSource(s string) bool {
	for _, source := range Sources {
		if string(source) == s {
			return true
		}
	}
	return false
}

type Article struct {
	URL    string
	Source Source
	Title  string
}
