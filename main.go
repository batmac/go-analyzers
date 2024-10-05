package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/appends"             // Package appends defines an Analyzer that detects if there is only one variable in append.
	"golang.org/x/tools/go/analysis/passes/asmdecl"             // Package asmdecl defines an Analyzer that reports mismatches between assembly files and Go declarations.
	"golang.org/x/tools/go/analysis/passes/assign"              // Package assign defines an Analyzer that detects useless assignments.
	"golang.org/x/tools/go/analysis/passes/atomic"              // Package atomic defines an Analyzer that checks for common mistakes using the sync/atomic package.
	"golang.org/x/tools/go/analysis/passes/atomicalign"         // Package atomicalign defines an Analyzer that checks for non-64-bit-aligned arguments to sync/atomic functions.
	"golang.org/x/tools/go/analysis/passes/bools"               // Package bools defines an Analyzer that detects common mistakes involving boolean operators.
	"golang.org/x/tools/go/analysis/passes/buildssa"            // Package buildssa defines an Analyzer that constructs the SSA representation of an error-free package and returns the set of all functions within it.
	"golang.org/x/tools/go/analysis/passes/buildtag"            // Package buildtag defines an Analyzer that checks build tags.
	"golang.org/x/tools/go/analysis/passes/cgocall"             // Package cgocall defines an Analyzer that detects some violations of the cgo pointer passing rules.
	"golang.org/x/tools/go/analysis/passes/composite"           // Package composite defines an Analyzer that checks for unkeyed composite literals.
	"golang.org/x/tools/go/analysis/passes/copylock"            // Package copylock defines an Analyzer that checks for locks erroneously passed by value.
	"golang.org/x/tools/go/analysis/passes/ctrlflow"            // Package ctrlflow is an analysis that provides a syntactic control-flow graph (CFG) for the body of a function.
	"golang.org/x/tools/go/analysis/passes/deepequalerrors"     // Package deepequalerrors defines an Analyzer that checks for the use of reflect.DeepEqual with error values.
	"golang.org/x/tools/go/analysis/passes/defers"              // Package defers defines an Analyzer that checks for common mistakes in defer statements.
	"golang.org/x/tools/go/analysis/passes/directive"           // Package directive defines an Analyzer that checks known Go toolchain directives.
	"golang.org/x/tools/go/analysis/passes/errorsas"            // The errorsas package defines an Analyzer that checks that the second argument to errors.As is a pointer to a type implementing error.
	"golang.org/x/tools/go/analysis/passes/fieldalignment"      // Package fieldalignment defines an Analyzer that detects structs that would use less memory if their fields were sorted.
	"golang.org/x/tools/go/analysis/passes/findcall"            // Package findcall defines an Analyzer that serves as a trivial example and test of the Analysis API.
	"golang.org/x/tools/go/analysis/passes/framepointer"        // Package framepointer defines an Analyzer that reports assembly code that clobbers the frame pointer before saving it.
	"golang.org/x/tools/go/analysis/passes/httpmux"             //
	"golang.org/x/tools/go/analysis/passes/httpresponse"        // Package httpresponse defines an Analyzer that checks for mistakes using HTTP responses.
	"golang.org/x/tools/go/analysis/passes/ifaceassert"         // Package ifaceassert defines an Analyzer that flags impossible interface-interface type assertions.
	"golang.org/x/tools/go/analysis/passes/inspect"             // Package inspect defines an Analyzer that provides an AST inspector (golang.org/x/tools/go/ast/inspector.Inspector) for the syntax trees of a package.
	"golang.org/x/tools/go/analysis/passes/loopclosure"         // Package loopclosure defines an Analyzer that checks for references to enclosing loop variables from within nested functions.
	"golang.org/x/tools/go/analysis/passes/lostcancel"          // Package lostcancel defines an Analyzer that checks for failure to call a context cancellation function.
	"golang.org/x/tools/go/analysis/passes/nilfunc"             // Package nilfunc defines an Analyzer that checks for useless comparisons against nil.
	"golang.org/x/tools/go/analysis/passes/nilness"             // Package nilness inspects the control-flow graph of an SSA function and reports errors such as nil pointer dereferences and degenerate nil pointer comparisons.
	"golang.org/x/tools/go/analysis/passes/pkgfact"             // The pkgfact package is a demonstration and test of the package fact mechanism.
	"golang.org/x/tools/go/analysis/passes/printf"              // Package printf defines an Analyzer that checks consistency of Printf format strings and arguments.
	"golang.org/x/tools/go/analysis/passes/reflectvaluecompare" // Package reflectvaluecompare defines an Analyzer that checks for accidentally using == or reflect.DeepEqual to compare reflect.Value values.
	"golang.org/x/tools/go/analysis/passes/shadow"              // Package shadow defines an Analyzer that checks for shadowed variables.
	"golang.org/x/tools/go/analysis/passes/shift"               // Package shift defines an Analyzer that checks for shifts that exceed the width of an integer.
	"golang.org/x/tools/go/analysis/passes/sigchanyzer"         // Package sigchanyzer defines an Analyzer that detects misuse of unbuffered signal as argument to signal.Notify.
	"golang.org/x/tools/go/analysis/passes/slog"                // Package slog defines an Analyzer that checks for mismatched key-value pairs in log/slog calls.
	"golang.org/x/tools/go/analysis/passes/sortslice"           // Package sortslice defines an Analyzer that checks for calls to sort.Slice that do not use a slice type as first argument.
	"golang.org/x/tools/go/analysis/passes/stdmethods"          // Package stdmethods defines an Analyzer that checks for misspellings in the signatures of methods similar to well-known interfaces.
	"golang.org/x/tools/go/analysis/passes/stdversion"          // Package stdversion reports uses of standard library symbols that are "too new" for the Go version in force in the referring file.
	"golang.org/x/tools/go/analysis/passes/stringintconv"       // Package stringintconv defines an Analyzer that flags type conversions from integers to strings.
	"golang.org/x/tools/go/analysis/passes/structtag"           // Package structtag defines an Analyzer that checks struct field tags are well formed.
	"golang.org/x/tools/go/analysis/passes/testinggoroutine"    // Package testinggoroutine defines an Analyzerfor detecting calls to Fatal from a test goroutine.
	"golang.org/x/tools/go/analysis/passes/tests"               // Package tests defines an Analyzer that checks for common mistaken usages of tests and examples.
	"golang.org/x/tools/go/analysis/passes/timeformat"          // Package timeformat defines an Analyzer that checks for the use of time.Format or time.Parse calls with a bad format.
	"golang.org/x/tools/go/analysis/passes/unmarshal"           // The unmarshal package defines an Analyzer that checks for passing non-pointer or non-interface types to unmarshal and decode functions.
	"golang.org/x/tools/go/analysis/passes/unreachable"         // Package unreachable defines an Analyzer that checks for unreachable code.
	"golang.org/x/tools/go/analysis/passes/unsafeptr"           // Package unsafeptr defines an Analyzer that checks for invalid conversions of uintptr to unsafe.Pointer.
	"golang.org/x/tools/go/analysis/passes/unusedresult"        // Package unusedresult defines an analyzer that checks for unused results of calls to certain pure functions.
	"golang.org/x/tools/go/analysis/passes/unusedwrite"         // Package unusedwrite checks for unused writes to the elements of a struct or array object.
	"golang.org/x/tools/go/analysis/passes/usesgenerics"        // Package usesgenerics defines an Analyzer that checks for usage of generic features added in Go 1.18.
)

func main() {
	multichecker.Main(
		appends.Analyzer,
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		atomicalign.Analyzer,
		bools.Analyzer,
		buildssa.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		ctrlflow.Analyzer,
		deepequalerrors.Analyzer,
		defers.Analyzer,
		directive.Analyzer,
		errorsas.Analyzer,
		fieldalignment.Analyzer,
		findcall.Analyzer,
		framepointer.Analyzer,
		httpmux.Analyzer,
		httpresponse.Analyzer,
		ifaceassert.Analyzer,
		inspect.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		nilness.Analyzer,
		pkgfact.Analyzer,
		printf.Analyzer,
		reflectvaluecompare.Analyzer,
		shadow.Analyzer,
		shift.Analyzer,
		sigchanyzer.Analyzer,
		slog.Analyzer,
		sortslice.Analyzer,
		stdmethods.Analyzer,
		stdversion.Analyzer,
		stringintconv.Analyzer,
		structtag.Analyzer,
		testinggoroutine.Analyzer,
		tests.Analyzer,
		timeformat.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
		unusedresult.Analyzer,
		unusedwrite.Analyzer,
		usesgenerics.Analyzer,
	)
}
