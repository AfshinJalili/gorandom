import type { Article } from "./types";

export const articles: Article[] = [
  // ============================================
  // Go Docs - Official Documentation
  // ============================================
  {
    url: "https://go.dev/doc/effective_go",
    source: "docs",
    title: "Effective Go",
  },
  {
    url: "https://go.dev/doc/faq",
    source: "docs",
    title: "Frequently Asked Questions",
  },
  {
    url: "https://go.dev/doc/code",
    source: "docs",
    title: "How to Write Go Code",
  },
  {
    url: "https://go.dev/doc/tutorial/getting-started",
    source: "docs",
    title: "Tutorial: Getting Started",
  },
  {
    url: "https://go.dev/doc/tutorial/create-module",
    source: "docs",
    title: "Tutorial: Create a Module",
  },
  {
    url: "https://go.dev/doc/tutorial/call-module-code",
    source: "docs",
    title: "Tutorial: Call Your Code From Another Module",
  },
  {
    url: "https://go.dev/doc/tutorial/handle-errors",
    source: "docs",
    title: "Tutorial: Return and Handle an Error",
  },
  {
    url: "https://go.dev/doc/tutorial/random-greeting",
    source: "docs",
    title: "Tutorial: Return a Random Greeting",
  },
  {
    url: "https://go.dev/doc/tutorial/greetings-multiple-people",
    source: "docs",
    title: "Tutorial: Return Greetings for Multiple People",
  },
  {
    url: "https://go.dev/doc/tutorial/add-a-test",
    source: "docs",
    title: "Tutorial: Add a Test",
  },
  {
    url: "https://go.dev/doc/tutorial/compile-install",
    source: "docs",
    title: "Tutorial: Compile and Install",
  },
  {
    url: "https://go.dev/doc/tutorial/web-service-gin",
    source: "docs",
    title: "Tutorial: Developing a RESTful API with Go and Gin",
  },
  {
    url: "https://go.dev/doc/tutorial/generics",
    source: "docs",
    title: "Tutorial: Getting Started with Generics",
  },
  {
    url: "https://go.dev/doc/tutorial/fuzz",
    source: "docs",
    title: "Tutorial: Getting Started with Fuzzing",
  },
  {
    url: "https://go.dev/doc/tutorial/workspaces",
    source: "docs",
    title: "Tutorial: Getting Started with Multi-module Workspaces",
  },
  {
    url: "https://go.dev/doc/tutorial/database-access",
    source: "docs",
    title: "Tutorial: Accessing a Relational Database",
  },
  {
    url: "https://go.dev/doc/articles/wiki/",
    source: "docs",
    title: "Writing Web Applications",
  },
  {
    url: "https://go.dev/doc/diagnostics",
    source: "docs",
    title: "Diagnostics",
  },
  {
    url: "https://go.dev/doc/gc-guide",
    source: "docs",
    title: "A Guide to the Go Garbage Collector",
  },
  {
    url: "https://go.dev/doc/pgo",
    source: "docs",
    title: "Profile-guided Optimization",
  },
  {
    url: "https://go.dev/doc/modules/managing-dependencies",
    source: "docs",
    title: "Managing Dependencies",
  },
  {
    url: "https://go.dev/doc/modules/developing",
    source: "docs",
    title: "Developing and Publishing Modules",
  },
  {
    url: "https://go.dev/doc/modules/layout",
    source: "docs",
    title: "Module Layout",
  },
  {
    url: "https://go.dev/doc/modules/version-numbers",
    source: "docs",
    title: "Module Version Numbers",
  },
  {
    url: "https://go.dev/ref/spec",
    source: "docs",
    title: "The Go Programming Language Specification",
  },
  {
    url: "https://go.dev/ref/mem",
    source: "docs",
    title: "The Go Memory Model",
  },

  // ============================================
  // Tour of Go - Interactive Tutorial
  // ============================================
  // Welcome
  {
    url: "https://go.dev/tour/welcome/1",
    source: "tour",
    title: "Hello, World",
  },
  { url: "https://go.dev/tour/welcome/2", source: "tour", title: "Go Local" },
  {
    url: "https://go.dev/tour/welcome/3",
    source: "tour",
    title: "The Go Playground",
  },
  {
    url: "https://go.dev/tour/welcome/4",
    source: "tour",
    title: "Congratulations",
  },
  // Basics
  { url: "https://go.dev/tour/basics/1", source: "tour", title: "Packages" },
  { url: "https://go.dev/tour/basics/2", source: "tour", title: "Imports" },
  {
    url: "https://go.dev/tour/basics/3",
    source: "tour",
    title: "Exported Names",
  },
  { url: "https://go.dev/tour/basics/4", source: "tour", title: "Functions" },
  {
    url: "https://go.dev/tour/basics/5",
    source: "tour",
    title: "Functions Continued",
  },
  {
    url: "https://go.dev/tour/basics/6",
    source: "tour",
    title: "Multiple Results",
  },
  {
    url: "https://go.dev/tour/basics/7",
    source: "tour",
    title: "Named Return Values",
  },
  { url: "https://go.dev/tour/basics/8", source: "tour", title: "Variables" },
  {
    url: "https://go.dev/tour/basics/9",
    source: "tour",
    title: "Variables with Initializers",
  },
  {
    url: "https://go.dev/tour/basics/10",
    source: "tour",
    title: "Short Variable Declarations",
  },
  {
    url: "https://go.dev/tour/basics/11",
    source: "tour",
    title: "Basic Types",
  },
  {
    url: "https://go.dev/tour/basics/12",
    source: "tour",
    title: "Zero Values",
  },
  {
    url: "https://go.dev/tour/basics/13",
    source: "tour",
    title: "Type Conversions",
  },
  {
    url: "https://go.dev/tour/basics/14",
    source: "tour",
    title: "Type Inference",
  },
  { url: "https://go.dev/tour/basics/15", source: "tour", title: "Constants" },
  {
    url: "https://go.dev/tour/basics/16",
    source: "tour",
    title: "Numeric Constants",
  },
  // Flow Control
  { url: "https://go.dev/tour/flowcontrol/1", source: "tour", title: "For" },
  {
    url: "https://go.dev/tour/flowcontrol/2",
    source: "tour",
    title: "For Continued",
  },
  {
    url: "https://go.dev/tour/flowcontrol/3",
    source: "tour",
    title: "For is Go's While",
  },
  {
    url: "https://go.dev/tour/flowcontrol/4",
    source: "tour",
    title: "Forever",
  },
  { url: "https://go.dev/tour/flowcontrol/5", source: "tour", title: "If" },
  {
    url: "https://go.dev/tour/flowcontrol/6",
    source: "tour",
    title: "If with Short Statement",
  },
  {
    url: "https://go.dev/tour/flowcontrol/7",
    source: "tour",
    title: "If and Else",
  },
  { url: "https://go.dev/tour/flowcontrol/8", source: "tour", title: "Switch" },
  {
    url: "https://go.dev/tour/flowcontrol/9",
    source: "tour",
    title: "Switch Evaluation Order",
  },
  {
    url: "https://go.dev/tour/flowcontrol/10",
    source: "tour",
    title: "Switch with No Condition",
  },
  { url: "https://go.dev/tour/flowcontrol/11", source: "tour", title: "Defer" },
  {
    url: "https://go.dev/tour/flowcontrol/12",
    source: "tour",
    title: "Stacking Defers",
  },
  // More Types
  { url: "https://go.dev/tour/moretypes/1", source: "tour", title: "Pointers" },
  { url: "https://go.dev/tour/moretypes/2", source: "tour", title: "Structs" },
  {
    url: "https://go.dev/tour/moretypes/3",
    source: "tour",
    title: "Struct Fields",
  },
  {
    url: "https://go.dev/tour/moretypes/4",
    source: "tour",
    title: "Pointers to Structs",
  },
  {
    url: "https://go.dev/tour/moretypes/5",
    source: "tour",
    title: "Struct Literals",
  },
  { url: "https://go.dev/tour/moretypes/6", source: "tour", title: "Arrays" },
  { url: "https://go.dev/tour/moretypes/7", source: "tour", title: "Slices" },
  {
    url: "https://go.dev/tour/moretypes/8",
    source: "tour",
    title: "Slices are Like References",
  },
  {
    url: "https://go.dev/tour/moretypes/9",
    source: "tour",
    title: "Slice Literals",
  },
  {
    url: "https://go.dev/tour/moretypes/10",
    source: "tour",
    title: "Slice Defaults",
  },
  {
    url: "https://go.dev/tour/moretypes/11",
    source: "tour",
    title: "Slice Length and Capacity",
  },
  {
    url: "https://go.dev/tour/moretypes/12",
    source: "tour",
    title: "Nil Slices",
  },
  {
    url: "https://go.dev/tour/moretypes/13",
    source: "tour",
    title: "Creating a Slice with Make",
  },
  {
    url: "https://go.dev/tour/moretypes/14",
    source: "tour",
    title: "Slices of Slices",
  },
  {
    url: "https://go.dev/tour/moretypes/15",
    source: "tour",
    title: "Appending to a Slice",
  },
  { url: "https://go.dev/tour/moretypes/16", source: "tour", title: "Range" },
  {
    url: "https://go.dev/tour/moretypes/17",
    source: "tour",
    title: "Range Continued",
  },
  { url: "https://go.dev/tour/moretypes/18", source: "tour", title: "Maps" },
  {
    url: "https://go.dev/tour/moretypes/19",
    source: "tour",
    title: "Map Literals",
  },
  {
    url: "https://go.dev/tour/moretypes/20",
    source: "tour",
    title: "Mutating Maps",
  },
  {
    url: "https://go.dev/tour/moretypes/21",
    source: "tour",
    title: "Function Values",
  },
  {
    url: "https://go.dev/tour/moretypes/22",
    source: "tour",
    title: "Function Closures",
  },
  // Methods and Interfaces
  { url: "https://go.dev/tour/methods/1", source: "tour", title: "Methods" },
  {
    url: "https://go.dev/tour/methods/2",
    source: "tour",
    title: "Methods are Functions",
  },
  {
    url: "https://go.dev/tour/methods/3",
    source: "tour",
    title: "Methods Continued",
  },
  {
    url: "https://go.dev/tour/methods/4",
    source: "tour",
    title: "Pointer Receivers",
  },
  {
    url: "https://go.dev/tour/methods/5",
    source: "tour",
    title: "Pointers and Functions",
  },
  {
    url: "https://go.dev/tour/methods/6",
    source: "tour",
    title: "Methods and Pointer Indirection",
  },
  {
    url: "https://go.dev/tour/methods/7",
    source: "tour",
    title: "Methods and Pointer Indirection (2)",
  },
  {
    url: "https://go.dev/tour/methods/8",
    source: "tour",
    title: "Choosing a Value or Pointer Receiver",
  },
  { url: "https://go.dev/tour/methods/9", source: "tour", title: "Interfaces" },
  {
    url: "https://go.dev/tour/methods/10",
    source: "tour",
    title: "Interfaces are Implemented Implicitly",
  },
  {
    url: "https://go.dev/tour/methods/11",
    source: "tour",
    title: "Interface Values",
  },
  {
    url: "https://go.dev/tour/methods/12",
    source: "tour",
    title: "Interface Values with Nil",
  },
  {
    url: "https://go.dev/tour/methods/13",
    source: "tour",
    title: "Nil Interface Values",
  },
  {
    url: "https://go.dev/tour/methods/14",
    source: "tour",
    title: "The Empty Interface",
  },
  {
    url: "https://go.dev/tour/methods/15",
    source: "tour",
    title: "Type Assertions",
  },
  {
    url: "https://go.dev/tour/methods/16",
    source: "tour",
    title: "Type Switches",
  },
  { url: "https://go.dev/tour/methods/17", source: "tour", title: "Stringers" },
  { url: "https://go.dev/tour/methods/18", source: "tour", title: "Errors" },
  { url: "https://go.dev/tour/methods/19", source: "tour", title: "Readers" },
  { url: "https://go.dev/tour/methods/20", source: "tour", title: "Images" },
  // Generics
  {
    url: "https://go.dev/tour/generics/1",
    source: "tour",
    title: "Type Parameters",
  },
  {
    url: "https://go.dev/tour/generics/2",
    source: "tour",
    title: "Generic Types",
  },
  // Concurrency
  {
    url: "https://go.dev/tour/concurrency/1",
    source: "tour",
    title: "Goroutines",
  },
  {
    url: "https://go.dev/tour/concurrency/2",
    source: "tour",
    title: "Channels",
  },
  {
    url: "https://go.dev/tour/concurrency/3",
    source: "tour",
    title: "Buffered Channels",
  },
  {
    url: "https://go.dev/tour/concurrency/4",
    source: "tour",
    title: "Range and Close",
  },
  { url: "https://go.dev/tour/concurrency/5", source: "tour", title: "Select" },
  {
    url: "https://go.dev/tour/concurrency/6",
    source: "tour",
    title: "Default Selection",
  },
  {
    url: "https://go.dev/tour/concurrency/7",
    source: "tour",
    title: "sync.Mutex",
  },

  // ============================================
  // Go by Example
  // ============================================
  {
    url: "https://gobyexample.com/hello-world",
    source: "gobyexample",
    title: "Hello World",
  },
  {
    url: "https://gobyexample.com/values",
    source: "gobyexample",
    title: "Values",
  },
  {
    url: "https://gobyexample.com/variables",
    source: "gobyexample",
    title: "Variables",
  },
  {
    url: "https://gobyexample.com/constants",
    source: "gobyexample",
    title: "Constants",
  },
  { url: "https://gobyexample.com/for", source: "gobyexample", title: "For" },
  {
    url: "https://gobyexample.com/if-else",
    source: "gobyexample",
    title: "If/Else",
  },
  {
    url: "https://gobyexample.com/switch",
    source: "gobyexample",
    title: "Switch",
  },
  {
    url: "https://gobyexample.com/arrays",
    source: "gobyexample",
    title: "Arrays",
  },
  {
    url: "https://gobyexample.com/slices",
    source: "gobyexample",
    title: "Slices",
  },
  { url: "https://gobyexample.com/maps", source: "gobyexample", title: "Maps" },
  {
    url: "https://gobyexample.com/functions",
    source: "gobyexample",
    title: "Functions",
  },
  {
    url: "https://gobyexample.com/multiple-return-values",
    source: "gobyexample",
    title: "Multiple Return Values",
  },
  {
    url: "https://gobyexample.com/variadic-functions",
    source: "gobyexample",
    title: "Variadic Functions",
  },
  {
    url: "https://gobyexample.com/closures",
    source: "gobyexample",
    title: "Closures",
  },
  {
    url: "https://gobyexample.com/recursion",
    source: "gobyexample",
    title: "Recursion",
  },
  {
    url: "https://gobyexample.com/pointers",
    source: "gobyexample",
    title: "Pointers",
  },
  {
    url: "https://gobyexample.com/strings-and-runes",
    source: "gobyexample",
    title: "Strings and Runes",
  },
  {
    url: "https://gobyexample.com/structs",
    source: "gobyexample",
    title: "Structs",
  },
  {
    url: "https://gobyexample.com/methods",
    source: "gobyexample",
    title: "Methods",
  },
  {
    url: "https://gobyexample.com/interfaces",
    source: "gobyexample",
    title: "Interfaces",
  },
  {
    url: "https://gobyexample.com/enums",
    source: "gobyexample",
    title: "Enums",
  },
  {
    url: "https://gobyexample.com/struct-embedding",
    source: "gobyexample",
    title: "Struct Embedding",
  },
  {
    url: "https://gobyexample.com/generics",
    source: "gobyexample",
    title: "Generics",
  },
  {
    url: "https://gobyexample.com/errors",
    source: "gobyexample",
    title: "Errors",
  },
  {
    url: "https://gobyexample.com/custom-errors",
    source: "gobyexample",
    title: "Custom Errors",
  },
  {
    url: "https://gobyexample.com/goroutines",
    source: "gobyexample",
    title: "Goroutines",
  },
  {
    url: "https://gobyexample.com/channels",
    source: "gobyexample",
    title: "Channels",
  },
  {
    url: "https://gobyexample.com/channel-buffering",
    source: "gobyexample",
    title: "Channel Buffering",
  },
  {
    url: "https://gobyexample.com/channel-synchronization",
    source: "gobyexample",
    title: "Channel Synchronization",
  },
  {
    url: "https://gobyexample.com/channel-directions",
    source: "gobyexample",
    title: "Channel Directions",
  },
  {
    url: "https://gobyexample.com/select",
    source: "gobyexample",
    title: "Select",
  },
  {
    url: "https://gobyexample.com/timeouts",
    source: "gobyexample",
    title: "Timeouts",
  },
  {
    url: "https://gobyexample.com/non-blocking-channel-operations",
    source: "gobyexample",
    title: "Non-Blocking Channel Operations",
  },
  {
    url: "https://gobyexample.com/closing-channels",
    source: "gobyexample",
    title: "Closing Channels",
  },
  {
    url: "https://gobyexample.com/range-over-channels",
    source: "gobyexample",
    title: "Range over Channels",
  },
  {
    url: "https://gobyexample.com/timers",
    source: "gobyexample",
    title: "Timers",
  },
  {
    url: "https://gobyexample.com/tickers",
    source: "gobyexample",
    title: "Tickers",
  },
  {
    url: "https://gobyexample.com/worker-pools",
    source: "gobyexample",
    title: "Worker Pools",
  },
  {
    url: "https://gobyexample.com/waitgroups",
    source: "gobyexample",
    title: "WaitGroups",
  },
  {
    url: "https://gobyexample.com/rate-limiting",
    source: "gobyexample",
    title: "Rate Limiting",
  },
  {
    url: "https://gobyexample.com/atomic-counters",
    source: "gobyexample",
    title: "Atomic Counters",
  },
  {
    url: "https://gobyexample.com/mutexes",
    source: "gobyexample",
    title: "Mutexes",
  },
  {
    url: "https://gobyexample.com/stateful-goroutines",
    source: "gobyexample",
    title: "Stateful Goroutines",
  },
  {
    url: "https://gobyexample.com/sorting",
    source: "gobyexample",
    title: "Sorting",
  },
  {
    url: "https://gobyexample.com/sorting-by-functions",
    source: "gobyexample",
    title: "Sorting by Functions",
  },
  {
    url: "https://gobyexample.com/panic",
    source: "gobyexample",
    title: "Panic",
  },
  {
    url: "https://gobyexample.com/defer",
    source: "gobyexample",
    title: "Defer",
  },
  {
    url: "https://gobyexample.com/recover",
    source: "gobyexample",
    title: "Recover",
  },
  {
    url: "https://gobyexample.com/string-functions",
    source: "gobyexample",
    title: "String Functions",
  },
  {
    url: "https://gobyexample.com/string-formatting",
    source: "gobyexample",
    title: "String Formatting",
  },
  {
    url: "https://gobyexample.com/text-templates",
    source: "gobyexample",
    title: "Text Templates",
  },
  {
    url: "https://gobyexample.com/regular-expressions",
    source: "gobyexample",
    title: "Regular Expressions",
  },
  { url: "https://gobyexample.com/json", source: "gobyexample", title: "JSON" },
  { url: "https://gobyexample.com/xml", source: "gobyexample", title: "XML" },
  { url: "https://gobyexample.com/time", source: "gobyexample", title: "Time" },
  {
    url: "https://gobyexample.com/epoch",
    source: "gobyexample",
    title: "Epoch",
  },
  {
    url: "https://gobyexample.com/time-formatting-parsing",
    source: "gobyexample",
    title: "Time Formatting / Parsing",
  },
  {
    url: "https://gobyexample.com/random-numbers",
    source: "gobyexample",
    title: "Random Numbers",
  },
  {
    url: "https://gobyexample.com/number-parsing",
    source: "gobyexample",
    title: "Number Parsing",
  },
  {
    url: "https://gobyexample.com/url-parsing",
    source: "gobyexample",
    title: "URL Parsing",
  },
  {
    url: "https://gobyexample.com/sha256-hashes",
    source: "gobyexample",
    title: "SHA256 Hashes",
  },
  {
    url: "https://gobyexample.com/base64-encoding",
    source: "gobyexample",
    title: "Base64 Encoding",
  },
  {
    url: "https://gobyexample.com/reading-files",
    source: "gobyexample",
    title: "Reading Files",
  },
  {
    url: "https://gobyexample.com/writing-files",
    source: "gobyexample",
    title: "Writing Files",
  },
  {
    url: "https://gobyexample.com/line-filters",
    source: "gobyexample",
    title: "Line Filters",
  },
  {
    url: "https://gobyexample.com/file-paths",
    source: "gobyexample",
    title: "File Paths",
  },
  {
    url: "https://gobyexample.com/directories",
    source: "gobyexample",
    title: "Directories",
  },
  {
    url: "https://gobyexample.com/temporary-files-and-directories",
    source: "gobyexample",
    title: "Temporary Files and Directories",
  },
  {
    url: "https://gobyexample.com/embed-directive",
    source: "gobyexample",
    title: "Embed Directive",
  },
  {
    url: "https://gobyexample.com/testing-and-benchmarking",
    source: "gobyexample",
    title: "Testing and Benchmarking",
  },
  {
    url: "https://gobyexample.com/command-line-arguments",
    source: "gobyexample",
    title: "Command-Line Arguments",
  },
  {
    url: "https://gobyexample.com/command-line-flags",
    source: "gobyexample",
    title: "Command-Line Flags",
  },
  {
    url: "https://gobyexample.com/command-line-subcommands",
    source: "gobyexample",
    title: "Command-Line Subcommands",
  },
  {
    url: "https://gobyexample.com/environment-variables",
    source: "gobyexample",
    title: "Environment Variables",
  },
  {
    url: "https://gobyexample.com/logging",
    source: "gobyexample",
    title: "Logging",
  },
  {
    url: "https://gobyexample.com/http-client",
    source: "gobyexample",
    title: "HTTP Client",
  },
  {
    url: "https://gobyexample.com/http-server",
    source: "gobyexample",
    title: "HTTP Server",
  },
  {
    url: "https://gobyexample.com/context",
    source: "gobyexample",
    title: "Context",
  },
  {
    url: "https://gobyexample.com/spawning-processes",
    source: "gobyexample",
    title: "Spawning Processes",
  },
  {
    url: "https://gobyexample.com/execing-processes",
    source: "gobyexample",
    title: "Exec'ing Processes",
  },
  {
    url: "https://gobyexample.com/signals",
    source: "gobyexample",
    title: "Signals",
  },
  { url: "https://gobyexample.com/exit", source: "gobyexample", title: "Exit" },

  // ============================================
  // Go Standard Library Packages
  // ============================================
  { url: "https://pkg.go.dev/fmt", source: "pkg", title: "Package fmt" },
  {
    url: "https://pkg.go.dev/net/http",
    source: "pkg",
    title: "Package net/http",
  },
  {
    url: "https://pkg.go.dev/encoding/json",
    source: "pkg",
    title: "Package encoding/json",
  },
  {
    url: "https://pkg.go.dev/strings",
    source: "pkg",
    title: "Package strings",
  },
  { url: "https://pkg.go.dev/bytes", source: "pkg", title: "Package bytes" },
  { url: "https://pkg.go.dev/io", source: "pkg", title: "Package io" },
  { url: "https://pkg.go.dev/os", source: "pkg", title: "Package os" },
  {
    url: "https://pkg.go.dev/context",
    source: "pkg",
    title: "Package context",
  },
  { url: "https://pkg.go.dev/sync", source: "pkg", title: "Package sync" },
  { url: "https://pkg.go.dev/time", source: "pkg", title: "Package time" },
  { url: "https://pkg.go.dev/regexp", source: "pkg", title: "Package regexp" },
  { url: "https://pkg.go.dev/flag", source: "pkg", title: "Package flag" },
  { url: "https://pkg.go.dev/log", source: "pkg", title: "Package log" },
  {
    url: "https://pkg.go.dev/testing",
    source: "pkg",
    title: "Package testing",
  },
  {
    url: "https://pkg.go.dev/database/sql",
    source: "pkg",
    title: "Package database/sql",
  },
  { url: "https://pkg.go.dev/bufio", source: "pkg", title: "Package bufio" },
  {
    url: "https://pkg.go.dev/path/filepath",
    source: "pkg",
    title: "Package path/filepath",
  },
  { url: "https://pkg.go.dev/sort", source: "pkg", title: "Package sort" },
  { url: "https://pkg.go.dev/math", source: "pkg", title: "Package math" },
  {
    url: "https://pkg.go.dev/reflect",
    source: "pkg",
    title: "Package reflect",
  },
  {
    url: "https://pkg.go.dev/runtime",
    source: "pkg",
    title: "Package runtime",
  },
  { url: "https://pkg.go.dev/errors", source: "pkg", title: "Package errors" },
  {
    url: "https://pkg.go.dev/html/template",
    source: "pkg",
    title: "Package html/template",
  },
  {
    url: "https://pkg.go.dev/text/template",
    source: "pkg",
    title: "Package text/template",
  },

  // ============================================
  // Go Blog - Selected Articles
  // ============================================
  {
    url: "https://go.dev/blog/defer-panic-and-recover",
    source: "blog",
    title: "Defer, Panic, and Recover",
  },
  {
    url: "https://go.dev/blog/error-handling-and-go",
    source: "blog",
    title: "Error Handling and Go",
  },
  {
    url: "https://go.dev/blog/errors-are-values",
    source: "blog",
    title: "Errors are Values",
  },
  {
    url: "https://go.dev/blog/go1.13-errors",
    source: "blog",
    title: "Working with Errors in Go 1.13",
  },
  {
    url: "https://go.dev/blog/slices",
    source: "blog",
    title: "Arrays, Slices: The Mechanics of Append",
  },
  {
    url: "https://go.dev/blog/slices-intro",
    source: "blog",
    title: "Go Slices: Usage and Internals",
  },
  {
    url: "https://go.dev/blog/strings",
    source: "blog",
    title: "Strings, Bytes, Runes and Characters in Go",
  },
  {
    url: "https://go.dev/blog/maps",
    source: "blog",
    title: "Go Maps in Action",
  },
  { url: "https://go.dev/blog/constants", source: "blog", title: "Constants" },
  {
    url: "https://go.dev/blog/laws-of-reflection",
    source: "blog",
    title: "The Laws of Reflection",
  },
  { url: "https://go.dev/blog/json", source: "blog", title: "JSON and Go" },
  { url: "https://go.dev/blog/gob", source: "blog", title: "Gobs of Data" },
  {
    url: "https://go.dev/blog/pipelines",
    source: "blog",
    title: "Go Concurrency Patterns: Pipelines and Cancellation",
  },
  {
    url: "https://go.dev/blog/context",
    source: "blog",
    title: "Go Concurrency Patterns: Context",
  },
  {
    url: "https://go.dev/blog/race-detector",
    source: "blog",
    title: "Introducing the Go Race Detector",
  },
  {
    url: "https://go.dev/blog/subtests",
    source: "blog",
    title: "Using Subtests and Sub-benchmarks",
  },
  {
    url: "https://go.dev/blog/cover",
    source: "blog",
    title: "The Cover Story",
  },
  {
    url: "https://go.dev/blog/pprof",
    source: "blog",
    title: "Profiling Go Programs",
  },
  {
    url: "https://go.dev/blog/gofmt",
    source: "blog",
    title: "go fmt Your Code",
  },
  {
    url: "https://go.dev/blog/godoc",
    source: "blog",
    title: "Godoc: Documenting Go Code",
  },
  {
    url: "https://go.dev/blog/organizing-go-code",
    source: "blog",
    title: "Organizing Go Code",
  },
  {
    url: "https://go.dev/blog/package-names",
    source: "blog",
    title: "Package Names",
  },
  {
    url: "https://go.dev/blog/declaration-syntax",
    source: "blog",
    title: "Go's Declaration Syntax",
  },
  { url: "https://go.dev/blog/cgo", source: "blog", title: "C? Go? Cgo!" },
  {
    url: "https://go.dev/blog/generate",
    source: "blog",
    title: "Generating Code",
  },
  {
    url: "https://go.dev/blog/using-go-modules",
    source: "blog",
    title: "Using Go Modules",
  },
  {
    url: "https://go.dev/blog/migrating-to-go-modules",
    source: "blog",
    title: "Migrating to Go Modules",
  },
  {
    url: "https://go.dev/blog/publishing-go-modules",
    source: "blog",
    title: "Publishing Go Modules",
  },
  {
    url: "https://go.dev/blog/v2-go-modules",
    source: "blog",
    title: "Go Modules: v2 and Beyond",
  },
  {
    url: "https://go.dev/blog/intro-generics",
    source: "blog",
    title: "An Introduction to Generics",
  },
  {
    url: "https://go.dev/blog/when-generics",
    source: "blog",
    title: "When to Use Generics",
  },
  {
    url: "https://go.dev/blog/why-generics",
    source: "blog",
    title: "Why Generics?",
  },
  {
    url: "https://go.dev/blog/go1.18",
    source: "blog",
    title: "Go 1.18 is Released!",
  },
  {
    url: "https://go.dev/blog/go1.21",
    source: "blog",
    title: "Go 1.21 is Released!",
  },
  {
    url: "https://go.dev/blog/go1.22",
    source: "blog",
    title: "Go 1.22 is Released!",
  },
  {
    url: "https://go.dev/blog/range-functions",
    source: "blog",
    title: "Range Over Function Types",
  },
  {
    url: "https://go.dev/blog/slog",
    source: "blog",
    title: "Structured Logging with slog",
  },
  {
    url: "https://go.dev/blog/pgo",
    source: "blog",
    title: "Profile-guided Optimization in Go 1.21",
  },
  {
    url: "https://go.dev/blog/routing-enhancements",
    source: "blog",
    title: "Routing Enhancements for Go 1.22",
  },
  {
    url: "https://go.dev/blog/wasi",
    source: "blog",
    title: "WASI Support in Go",
  },
  {
    url: "https://go.dev/blog/govulncheck",
    source: "blog",
    title: "Govulncheck v1.0.0 is Released!",
  },
  {
    url: "https://go.dev/blog/go15gc",
    source: "blog",
    title: "Go GC: Prioritizing Low Latency and Simplicity",
  },
  {
    url: "https://go.dev/blog/ismmkeynote",
    source: "blog",
    title: "Getting to Go: The Journey of Go's Garbage Collector",
  },
];
