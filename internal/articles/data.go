package articles

var Data = []Article{
	{
		URL:    "https://go.dev/doc/effective_go",
		Source: SourceDocs,
		Title:  "Effective Go",
	},
	{
		URL:    "https://go.dev/doc/faq",
		Source: SourceDocs,
		Title:  "Frequently Asked Questions",
	},
	{
		URL:    "https://go.dev/doc/code",
		Source: SourceDocs,
		Title:  "How to Write Go Code",
	},
	{
		URL:    "https://go.dev/doc/tutorial/getting-started",
		Source: SourceDocs,
		Title:  "Tutorial: Getting Started",
	},
	{
		URL:    "https://go.dev/doc/tutorial/create-module",
		Source: SourceDocs,
		Title:  "Tutorial: Create a Module",
	},
	{
		URL:    "https://go.dev/doc/tutorial/call-module-code",
		Source: SourceDocs,
		Title:  "Tutorial: Call Your Code From Another Module",
	},
	{
		URL:    "https://go.dev/doc/tutorial/handle-errors",
		Source: SourceDocs,
		Title:  "Tutorial: Return and Handle an Error",
	},
	{
		URL:    "https://go.dev/doc/tutorial/random-greeting",
		Source: SourceDocs,
		Title:  "Tutorial: Return a Random Greeting",
	},
	{
		URL:    "https://go.dev/doc/tutorial/greetings-multiple-people",
		Source: SourceDocs,
		Title:  "Tutorial: Return Greetings for Multiple People",
	},
	{
		URL:    "https://go.dev/doc/tutorial/add-a-test",
		Source: SourceDocs,
		Title:  "Tutorial: Add a Test",
	},
	{
		URL:    "https://go.dev/doc/tutorial/compile-install",
		Source: SourceDocs,
		Title:  "Tutorial: Compile and Install",
	},
	{
		URL:    "https://go.dev/doc/tutorial/web-service-gin",
		Source: SourceDocs,
		Title:  "Tutorial: Developing a RESTful API with Go and Gin",
	},
	{
		URL:    "https://go.dev/doc/tutorial/generics",
		Source: SourceDocs,
		Title:  "Tutorial: Getting Started with Generics",
	},
	{
		URL:    "https://go.dev/doc/tutorial/fuzz",
		Source: SourceDocs,
		Title:  "Tutorial: Getting Started with Fuzzing",
	},
	{
		URL:    "https://go.dev/doc/tutorial/workspaces",
		Source: SourceDocs,
		Title:  "Tutorial: Getting Started with Multi-module Workspaces",
	},
	{
		URL:    "https://go.dev/doc/tutorial/database-access",
		Source: SourceDocs,
		Title:  "Tutorial: Accessing a Relational Database",
	},
	{
		URL:    "https://go.dev/doc/articles/wiki/",
		Source: SourceDocs,
		Title:  "Writing Web Applications",
	},
	{
		URL:    "https://go.dev/doc/diagnostics",
		Source: SourceDocs,
		Title:  "Diagnostics",
	},
	{
		URL:    "https://go.dev/doc/gc-guide",
		Source: SourceDocs,
		Title:  "A Guide to the Go Garbage Collector",
	},
	{
		URL:    "https://go.dev/doc/pgo",
		Source: SourceDocs,
		Title:  "Profile-guided Optimization",
	},
	{
		URL:    "https://go.dev/doc/modules/managing-dependencies",
		Source: SourceDocs,
		Title:  "Managing Dependencies",
	},
	{
		URL:    "https://go.dev/doc/modules/developing",
		Source: SourceDocs,
		Title:  "Developing and Publishing Modules",
	},
	{
		URL:    "https://go.dev/doc/modules/layout",
		Source: SourceDocs,
		Title:  "Module Layout",
	},
	{
		URL:    "https://go.dev/doc/modules/version-numbers",
		Source: SourceDocs,
		Title:  "Module Version Numbers",
	},
	{
		URL:    "https://go.dev/ref/spec",
		Source: SourceDocs,
		Title:  "The Go Programming Language Specification",
	},
	{
		URL:    "https://go.dev/ref/mem",
		Source: SourceDocs,
		Title:  "The Go Memory Model",
	},
	{
		URL:    "https://go.dev/tour/welcome/1",
		Source: SourceTour,
		Title:  "Hello, World",
	},
	{
		URL:    "https://go.dev/tour/welcome/3",
		Source: SourceTour,
		Title:  "The Go Playground",
	},
	{
		URL:    "https://go.dev/tour/welcome/4",
		Source: SourceTour,
		Title:  "Congratulations",
	},
	{
		URL:    "https://go.dev/tour/basics/3",
		Source: SourceTour,
		Title:  "Exported Names",
	},
	{
		URL:    "https://go.dev/tour/basics/5",
		Source: SourceTour,
		Title:  "Functions Continued",
	},
	{
		URL:    "https://go.dev/tour/basics/6",
		Source: SourceTour,
		Title:  "Multiple Results",
	},
	{
		URL:    "https://go.dev/tour/basics/7",
		Source: SourceTour,
		Title:  "Named Return Values",
	},
	{
		URL:    "https://go.dev/tour/basics/9",
		Source: SourceTour,
		Title:  "Variables with Initializers",
	},
	{
		URL:    "https://go.dev/tour/basics/10",
		Source: SourceTour,
		Title:  "Short Variable Declarations",
	},
	{
		URL:    "https://go.dev/tour/basics/11",
		Source: SourceTour,
		Title:  "Basic Types",
	},
	{
		URL:    "https://go.dev/tour/basics/12",
		Source: SourceTour,
		Title:  "Zero Values",
	},
	{
		URL:    "https://go.dev/tour/basics/13",
		Source: SourceTour,
		Title:  "Type Conversions",
	},
	{
		URL:    "https://go.dev/tour/basics/14",
		Source: SourceTour,
		Title:  "Type Inference",
	},
	{
		URL:    "https://go.dev/tour/basics/16",
		Source: SourceTour,
		Title:  "Numeric Constants",
	},
	{
		URL:    "https://go.dev/tour/flowcontrol/2",
		Source: SourceTour,
		Title:  "For Continued",
	},
	{
		URL:    "https://go.dev/tour/flowcontrol/3",
		Source: SourceTour,
		Title:  "For is Go's While",
	},
	{
		URL:    "https://go.dev/tour/flowcontrol/4",
		Source: SourceTour,
		Title:  "Forever",
	},
	{
		URL:    "https://go.dev/tour/flowcontrol/6",
		Source: SourceTour,
		Title:  "If with Short Statement",
	},
	{
		URL:    "https://go.dev/tour/flowcontrol/7",
		Source: SourceTour,
		Title:  "If and Else",
	},
	{
		URL:    "https://go.dev/tour/flowcontrol/9",
		Source: SourceTour,
		Title:  "Switch Evaluation Order",
	},
	{
		URL:    "https://go.dev/tour/flowcontrol/10",
		Source: SourceTour,
		Title:  "Switch with No Condition",
	},
	{
		URL:    "https://go.dev/tour/flowcontrol/12",
		Source: SourceTour,
		Title:  "Stacking Defers",
	},
	{
		URL:    "https://go.dev/tour/moretypes/3",
		Source: SourceTour,
		Title:  "Struct Fields",
	},
	{
		URL:    "https://go.dev/tour/moretypes/4",
		Source: SourceTour,
		Title:  "Pointers to Structs",
	},
	{
		URL:    "https://go.dev/tour/moretypes/5",
		Source: SourceTour,
		Title:  "Struct Literals",
	},
	{
		URL:    "https://go.dev/tour/moretypes/8",
		Source: SourceTour,
		Title:  "Slices are Like References",
	},
	{
		URL:    "https://go.dev/tour/moretypes/9",
		Source: SourceTour,
		Title:  "Slice Literals",
	},
	{
		URL:    "https://go.dev/tour/moretypes/10",
		Source: SourceTour,
		Title:  "Slice Defaults",
	},
	{
		URL:    "https://go.dev/tour/moretypes/11",
		Source: SourceTour,
		Title:  "Slice Length and Capacity",
	},
	{
		URL:    "https://go.dev/tour/moretypes/12",
		Source: SourceTour,
		Title:  "Nil Slices",
	},
	{
		URL:    "https://go.dev/tour/moretypes/13",
		Source: SourceTour,
		Title:  "Creating a Slice with Make",
	},
	{
		URL:    "https://go.dev/tour/moretypes/14",
		Source: SourceTour,
		Title:  "Slices of Slices",
	},
	{
		URL:    "https://go.dev/tour/moretypes/15",
		Source: SourceTour,
		Title:  "Appending to a Slice",
	},
	{
		URL:    "https://go.dev/tour/moretypes/17",
		Source: SourceTour,
		Title:  "Range Continued",
	},
	{
		URL:    "https://go.dev/tour/moretypes/19",
		Source: SourceTour,
		Title:  "Map Literals",
	},
	{
		URL:    "https://go.dev/tour/moretypes/20",
		Source: SourceTour,
		Title:  "Mutating Maps",
	},
	{
		URL:    "https://go.dev/tour/moretypes/21",
		Source: SourceTour,
		Title:  "Function Values",
	},
	{
		URL:    "https://go.dev/tour/moretypes/22",
		Source: SourceTour,
		Title:  "Function Closures",
	},
	{
		URL:    "https://go.dev/tour/methods/2",
		Source: SourceTour,
		Title:  "Methods are Functions",
	},
	{
		URL:    "https://go.dev/tour/methods/3",
		Source: SourceTour,
		Title:  "Methods Continued",
	},
	{
		URL:    "https://go.dev/tour/methods/4",
		Source: SourceTour,
		Title:  "Pointer Receivers",
	},
	{
		URL:    "https://go.dev/tour/methods/5",
		Source: SourceTour,
		Title:  "Pointers and Functions",
	},
	{
		URL:    "https://go.dev/tour/methods/6",
		Source: SourceTour,
		Title:  "Methods and Pointer Indirection",
	},
	{
		URL:    "https://go.dev/tour/methods/7",
		Source: SourceTour,
		Title:  "Methods and Pointer Indirection (2)",
	},
	{
		URL:    "https://go.dev/tour/methods/8",
		Source: SourceTour,
		Title:  "Choosing a Value or Pointer Receiver",
	},
	{
		URL:    "https://go.dev/tour/methods/10",
		Source: SourceTour,
		Title:  "Interfaces are Implemented Implicitly",
	},
	{
		URL:    "https://go.dev/tour/methods/11",
		Source: SourceTour,
		Title:  "Interface Values",
	},
	{
		URL:    "https://go.dev/tour/methods/12",
		Source: SourceTour,
		Title:  "Interface Values with Nil",
	},
	{
		URL:    "https://go.dev/tour/methods/13",
		Source: SourceTour,
		Title:  "Nil Interface Values",
	},
	{
		URL:    "https://go.dev/tour/methods/14",
		Source: SourceTour,
		Title:  "The Empty Interface",
	},
	{
		URL:    "https://go.dev/tour/methods/15",
		Source: SourceTour,
		Title:  "Type Assertions",
	},
	{
		URL:    "https://go.dev/tour/methods/16",
		Source: SourceTour,
		Title:  "Type Switches",
	},
	{
		URL:    "https://go.dev/tour/generics/1",
		Source: SourceTour,
		Title:  "Type Parameters",
	},
	{
		URL:    "https://go.dev/tour/generics/2",
		Source: SourceTour,
		Title:  "Generic Types",
	},
	{
		URL:    "https://go.dev/tour/concurrency/1",
		Source: SourceTour,
		Title:  "Goroutines",
	},
	{
		URL:    "https://go.dev/tour/concurrency/2",
		Source: SourceTour,
		Title:  "Channels",
	},
	{
		URL:    "https://go.dev/tour/concurrency/3",
		Source: SourceTour,
		Title:  "Buffered Channels",
	},
	{
		URL:    "https://go.dev/tour/concurrency/4",
		Source: SourceTour,
		Title:  "Range and Close",
	},
	{
		URL:    "https://go.dev/tour/concurrency/6",
		Source: SourceTour,
		Title:  "Default Selection",
	},
	{
		URL:    "https://go.dev/tour/concurrency/7",
		Source: SourceTour,
		Title:  "sync.Mutex",
	},
	{
		URL:    "https://gobyexample.com/hello-world",
		Source: SourceGoByExample,
		Title:  "Hello World",
	},
	{
		URL:    "https://gobyexample.com/values",
		Source: SourceGoByExample,
		Title:  "Values",
	},
	{
		URL:    "https://gobyexample.com/variables",
		Source: SourceGoByExample,
		Title:  "Variables",
	},
	{
		URL:    "https://gobyexample.com/constants",
		Source: SourceGoByExample,
		Title:  "Constants",
	},
	{
		URL:    "https://gobyexample.com/if-else",
		Source: SourceGoByExample,
		Title:  "If/Else",
	},
	{
		URL:    "https://gobyexample.com/switch",
		Source: SourceGoByExample,
		Title:  "Switch",
	},
	{
		URL:    "https://gobyexample.com/arrays",
		Source: SourceGoByExample,
		Title:  "Arrays",
	},
	{
		URL:    "https://gobyexample.com/slices",
		Source: SourceGoByExample,
		Title:  "Slices",
	},
	{
		URL:    "https://gobyexample.com/functions",
		Source: SourceGoByExample,
		Title:  "Functions",
	},
	{
		URL:    "https://gobyexample.com/multiple-return-values",
		Source: SourceGoByExample,
		Title:  "Multiple Return Values",
	},
	{
		URL:    "https://gobyexample.com/variadic-functions",
		Source: SourceGoByExample,
		Title:  "Variadic Functions",
	},
	{
		URL:    "https://gobyexample.com/closures",
		Source: SourceGoByExample,
		Title:  "Closures",
	},
	{
		URL:    "https://gobyexample.com/recursion",
		Source: SourceGoByExample,
		Title:  "Recursion",
	},
	{
		URL:    "https://gobyexample.com/pointers",
		Source: SourceGoByExample,
		Title:  "Pointers",
	},
	{
		URL:    "https://gobyexample.com/strings-and-runes",
		Source: SourceGoByExample,
		Title:  "Strings and Runes",
	},
	{
		URL:    "https://gobyexample.com/structs",
		Source: SourceGoByExample,
		Title:  "Structs",
	},
	{
		URL:    "https://gobyexample.com/methods",
		Source: SourceGoByExample,
		Title:  "Methods",
	},
	{
		URL:    "https://gobyexample.com/interfaces",
		Source: SourceGoByExample,
		Title:  "Interfaces",
	},
	{
		URL:    "https://gobyexample.com/enums",
		Source: SourceGoByExample,
		Title:  "Enums",
	},
	{
		URL:    "https://gobyexample.com/struct-embedding",
		Source: SourceGoByExample,
		Title:  "Struct Embedding",
	},
	{
		URL:    "https://gobyexample.com/generics",
		Source: SourceGoByExample,
		Title:  "Generics",
	},
	{
		URL:    "https://gobyexample.com/errors",
		Source: SourceGoByExample,
		Title:  "Errors",
	},
	{
		URL:    "https://gobyexample.com/custom-errors",
		Source: SourceGoByExample,
		Title:  "Custom Errors",
	},
	{
		URL:    "https://gobyexample.com/goroutines",
		Source: SourceGoByExample,
		Title:  "Goroutines",
	},
	{
		URL:    "https://gobyexample.com/channels",
		Source: SourceGoByExample,
		Title:  "Channels",
	},
	{
		URL:    "https://gobyexample.com/channel-buffering",
		Source: SourceGoByExample,
		Title:  "Channel Buffering",
	},
	{
		URL:    "https://gobyexample.com/channel-synchronization",
		Source: SourceGoByExample,
		Title:  "Channel Synchronization",
	},
	{
		URL:    "https://gobyexample.com/channel-directions",
		Source: SourceGoByExample,
		Title:  "Channel Directions",
	},
	{
		URL:    "https://gobyexample.com/select",
		Source: SourceGoByExample,
		Title:  "Select",
	},
	{
		URL:    "https://gobyexample.com/timeouts",
		Source: SourceGoByExample,
		Title:  "Timeouts",
	},
	{
		URL:    "https://gobyexample.com/non-blocking-channel-operations",
		Source: SourceGoByExample,
		Title:  "Non-Blocking Channel Operations",
	},
	{
		URL:    "https://gobyexample.com/closing-channels",
		Source: SourceGoByExample,
		Title:  "Closing Channels",
	},
	{
		URL:    "https://gobyexample.com/range-over-channels",
		Source: SourceGoByExample,
		Title:  "Range over Channels",
	},
	{
		URL:    "https://gobyexample.com/timers",
		Source: SourceGoByExample,
		Title:  "Timers",
	},
	{
		URL:    "https://gobyexample.com/tickers",
		Source: SourceGoByExample,
		Title:  "Tickers",
	},
	{
		URL:    "https://gobyexample.com/worker-pools",
		Source: SourceGoByExample,
		Title:  "Worker Pools",
	},
	{
		URL:    "https://gobyexample.com/waitgroups",
		Source: SourceGoByExample,
		Title:  "WaitGroups",
	},
	{
		URL:    "https://gobyexample.com/rate-limiting",
		Source: SourceGoByExample,
		Title:  "Rate Limiting",
	},
	{
		URL:    "https://gobyexample.com/atomic-counters",
		Source: SourceGoByExample,
		Title:  "Atomic Counters",
	},
	{
		URL:    "https://gobyexample.com/mutexes",
		Source: SourceGoByExample,
		Title:  "Mutexes",
	},
	{
		URL:    "https://gobyexample.com/stateful-goroutines",
		Source: SourceGoByExample,
		Title:  "Stateful Goroutines",
	},
	{
		URL:    "https://gobyexample.com/sorting",
		Source: SourceGoByExample,
		Title:  "Sorting",
	},
	{
		URL:    "https://gobyexample.com/sorting-by-functions",
		Source: SourceGoByExample,
		Title:  "Sorting by Functions",
	},
	{
		URL:    "https://gobyexample.com/panic",
		Source: SourceGoByExample,
		Title:  "Panic",
	},
	{
		URL:    "https://gobyexample.com/defer",
		Source: SourceGoByExample,
		Title:  "Defer",
	},
	{
		URL:    "https://gobyexample.com/recover",
		Source: SourceGoByExample,
		Title:  "Recover",
	},
	{
		URL:    "https://gobyexample.com/string-functions",
		Source: SourceGoByExample,
		Title:  "String Functions",
	},
	{
		URL:    "https://gobyexample.com/string-formatting",
		Source: SourceGoByExample,
		Title:  "String Formatting",
	},
	{
		URL:    "https://gobyexample.com/text-templates",
		Source: SourceGoByExample,
		Title:  "Text Templates",
	},
	{
		URL:    "https://gobyexample.com/regular-expressions",
		Source: SourceGoByExample,
		Title:  "Regular Expressions",
	},
	{
		URL:    "https://gobyexample.com/epoch",
		Source: SourceGoByExample,
		Title:  "Epoch",
	},
	{
		URL:    "https://gobyexample.com/time-formatting-parsing",
		Source: SourceGoByExample,
		Title:  "Time Formatting / Parsing",
	},
	{
		URL:    "https://gobyexample.com/random-numbers",
		Source: SourceGoByExample,
		Title:  "Random Numbers",
	},
	{
		URL:    "https://gobyexample.com/number-parsing",
		Source: SourceGoByExample,
		Title:  "Number Parsing",
	},
	{
		URL:    "https://gobyexample.com/url-parsing",
		Source: SourceGoByExample,
		Title:  "URL Parsing",
	},
	{
		URL:    "https://gobyexample.com/sha256-hashes",
		Source: SourceGoByExample,
		Title:  "SHA256 Hashes",
	},
	{
		URL:    "https://gobyexample.com/base64-encoding",
		Source: SourceGoByExample,
		Title:  "Base64 Encoding",
	},
	{
		URL:    "https://gobyexample.com/reading-files",
		Source: SourceGoByExample,
		Title:  "Reading Files",
	},
	{
		URL:    "https://gobyexample.com/writing-files",
		Source: SourceGoByExample,
		Title:  "Writing Files",
	},
	{
		URL:    "https://gobyexample.com/line-filters",
		Source: SourceGoByExample,
		Title:  "Line Filters",
	},
	{
		URL:    "https://gobyexample.com/file-paths",
		Source: SourceGoByExample,
		Title:  "File Paths",
	},
	{
		URL:    "https://gobyexample.com/directories",
		Source: SourceGoByExample,
		Title:  "Directories",
	},
	{
		URL:    "https://gobyexample.com/temporary-files-and-directories",
		Source: SourceGoByExample,
		Title:  "Temporary Files and Directories",
	},
	{
		URL:    "https://gobyexample.com/embed-directive",
		Source: SourceGoByExample,
		Title:  "Embed Directive",
	},
	{
		URL:    "https://gobyexample.com/testing-and-benchmarking",
		Source: SourceGoByExample,
		Title:  "Testing and Benchmarking",
	},
	{
		URL:    "https://gobyexample.com/command-line-arguments",
		Source: SourceGoByExample,
		Title:  "Command-Line Arguments",
	},
	{
		URL:    "https://gobyexample.com/command-line-flags",
		Source: SourceGoByExample,
		Title:  "Command-Line Flags",
	},
	{
		URL:    "https://gobyexample.com/command-line-subcommands",
		Source: SourceGoByExample,
		Title:  "Command-Line Subcommands",
	},
	{
		URL:    "https://gobyexample.com/environment-variables",
		Source: SourceGoByExample,
		Title:  "Environment Variables",
	},
	{
		URL:    "https://gobyexample.com/logging",
		Source: SourceGoByExample,
		Title:  "Logging",
	},
	{
		URL:    "https://gobyexample.com/http-client",
		Source: SourceGoByExample,
		Title:  "HTTP Client",
	},
	{
		URL:    "https://gobyexample.com/http-server",
		Source: SourceGoByExample,
		Title:  "HTTP Server",
	},
	{
		URL:    "https://gobyexample.com/context",
		Source: SourceGoByExample,
		Title:  "Context",
	},
	{
		URL:    "https://gobyexample.com/spawning-processes",
		Source: SourceGoByExample,
		Title:  "Spawning Processes",
	},
	{
		URL:    "https://gobyexample.com/execing-processes",
		Source: SourceGoByExample,
		Title:  "Exec'ing Processes",
	},
	{
		URL:    "https://gobyexample.com/signals",
		Source: SourceGoByExample,
		Title:  "Signals",
	},
	{
		URL:    "https://pkg.go.dev/net/http",
		Source: SourcePkg,
		Title:  "Package net/http",
	},
	{
		URL:    "https://pkg.go.dev/encoding/json",
		Source: SourcePkg,
		Title:  "Package encoding/json",
	},
	{
		URL:    "https://pkg.go.dev/strings",
		Source: SourcePkg,
		Title:  "Package strings",
	},
	{
		URL:    "https://pkg.go.dev/context",
		Source: SourcePkg,
		Title:  "Package context",
	},
	{
		URL:    "https://pkg.go.dev/testing",
		Source: SourcePkg,
		Title:  "Package testing",
	},
	{
		URL:    "https://pkg.go.dev/database/sql",
		Source: SourcePkg,
		Title:  "Package database/sql",
	},
	{
		URL:    "https://pkg.go.dev/path/filepath",
		Source: SourcePkg,
		Title:  "Package path/filepath",
	},
	{
		URL:    "https://pkg.go.dev/reflect",
		Source: SourcePkg,
		Title:  "Package reflect",
	},
	{
		URL:    "https://pkg.go.dev/runtime",
		Source: SourcePkg,
		Title:  "Package runtime",
	},
	{
		URL:    "https://pkg.go.dev/html/template",
		Source: SourcePkg,
		Title:  "Package html/template",
	},
	{
		URL:    "https://pkg.go.dev/text/template",
		Source: SourcePkg,
		Title:  "Package text/template",
	},
	{
		URL:    "https://go.dev/blog/defer-panic-and-recover",
		Source: SourceBlog,
		Title:  "Defer, Panic, and Recover",
	},
	{
		URL:    "https://go.dev/blog/error-handling-and-go",
		Source: SourceBlog,
		Title:  "Error Handling and Go",
	},
	{
		URL:    "https://go.dev/blog/errors-are-values",
		Source: SourceBlog,
		Title:  "Errors are Values",
	},
	{
		URL:    "https://go.dev/blog/go1.13-errors",
		Source: SourceBlog,
		Title:  "Working with Errors in Go 1.13",
	},
	{
		URL:    "https://go.dev/blog/slices",
		Source: SourceBlog,
		Title:  "Arrays, Slices: The Mechanics of Append",
	},
	{
		URL:    "https://go.dev/blog/slices-intro",
		Source: SourceBlog,
		Title:  "Go Slices: Usage and Internals",
	},
	{
		URL:    "https://go.dev/blog/strings",
		Source: SourceBlog,
		Title:  "Strings, Bytes, Runes and Characters in Go",
	},
	{
		URL:    "https://go.dev/blog/maps",
		Source: SourceBlog,
		Title:  "Go Maps in Action",
	},
	{
		URL:    "https://go.dev/blog/laws-of-reflection",
		Source: SourceBlog,
		Title:  "The Laws of Reflection",
	},
	{
		URL:    "https://go.dev/blog/pipelines",
		Source: SourceBlog,
		Title:  "Go Concurrency Patterns: Pipelines and Cancellation",
	},
	{
		URL:    "https://go.dev/blog/context",
		Source: SourceBlog,
		Title:  "Go Concurrency Patterns: Context",
	},
	{
		URL:    "https://go.dev/blog/race-detector",
		Source: SourceBlog,
		Title:  "Introducing the Go Race Detector",
	},
	{
		URL:    "https://go.dev/blog/subtests",
		Source: SourceBlog,
		Title:  "Using Subtests and Sub-benchmarks",
	},
	{
		URL:    "https://go.dev/blog/cover",
		Source: SourceBlog,
		Title:  "The Cover Story",
	},
	{
		URL:    "https://go.dev/blog/pprof",
		Source: SourceBlog,
		Title:  "Profiling Go Programs",
	},
	{
		URL:    "https://go.dev/blog/gofmt",
		Source: SourceBlog,
		Title:  "go fmt Your Code",
	},
	{
		URL:    "https://go.dev/blog/godoc",
		Source: SourceBlog,
		Title:  "Godoc: Documenting Go Code",
	},
	{
		URL:    "https://go.dev/blog/organizing-go-code",
		Source: SourceBlog,
		Title:  "Organizing Go Code",
	},
	{
		URL:    "https://go.dev/blog/package-names",
		Source: SourceBlog,
		Title:  "Package Names",
	},
	{
		URL:    "https://go.dev/blog/declaration-syntax",
		Source: SourceBlog,
		Title:  "Go's Declaration Syntax",
	},
	{
		URL:    "https://go.dev/blog/generate",
		Source: SourceBlog,
		Title:  "Generating Code",
	},
	{
		URL:    "https://go.dev/blog/using-go-modules",
		Source: SourceBlog,
		Title:  "Using Go Modules",
	},
	{
		URL:    "https://go.dev/blog/migrating-to-go-modules",
		Source: SourceBlog,
		Title:  "Migrating to Go Modules",
	},
	{
		URL:    "https://go.dev/blog/publishing-go-modules",
		Source: SourceBlog,
		Title:  "Publishing Go Modules",
	},
	{
		URL:    "https://go.dev/blog/v2-go-modules",
		Source: SourceBlog,
		Title:  "Go Modules: v2 and Beyond",
	},
	{
		URL:    "https://go.dev/blog/intro-generics",
		Source: SourceBlog,
		Title:  "An Introduction to Generics",
	},
	{
		URL:    "https://go.dev/blog/when-generics",
		Source: SourceBlog,
		Title:  "When to Use Generics",
	},
	{
		URL:    "https://go.dev/blog/why-generics",
		Source: SourceBlog,
		Title:  "Why Generics?",
	},
	{
		URL:    "https://go.dev/blog/go1.18",
		Source: SourceBlog,
		Title:  "Go 1.18 is Released!",
	},
	{
		URL:    "https://go.dev/blog/go1.21",
		Source: SourceBlog,
		Title:  "Go 1.21 is Released!",
	},
	{
		URL:    "https://go.dev/blog/go1.22",
		Source: SourceBlog,
		Title:  "Go 1.22 is Released!",
	},
	{
		URL:    "https://go.dev/blog/range-functions",
		Source: SourceBlog,
		Title:  "Range Over Function Types",
	},
	{
		URL:    "https://go.dev/blog/slog",
		Source: SourceBlog,
		Title:  "Structured Logging with slog",
	},
	{
		URL:    "https://go.dev/blog/pgo",
		Source: SourceBlog,
		Title:  "Profile-guided Optimization in Go 1.21",
	},
	{
		URL:    "https://go.dev/blog/routing-enhancements",
		Source: SourceBlog,
		Title:  "Routing Enhancements for Go 1.22",
	},
	{
		URL:    "https://go.dev/blog/wasi",
		Source: SourceBlog,
		Title:  "WASI Support in Go",
	},
	{
		URL:    "https://go.dev/blog/govulncheck",
		Source: SourceBlog,
		Title:  "Govulncheck v1.0.0 is Released!",
	},
	{
		URL:    "https://go.dev/blog/go15gc",
		Source: SourceBlog,
		Title:  "Go GC: Prioritizing Low Latency and Simplicity",
	},
	{
		URL:    "https://go.dev/blog/ismmkeynote",
		Source: SourceBlog,
		Title:  "Getting to Go: The Journey of Go's Garbage Collector",
	},
}