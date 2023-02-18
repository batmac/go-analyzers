package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/asmdecl"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/atomic"
	"golang.org/x/tools/go/analysis/passes/atomicalign"
	"golang.org/x/tools/go/analysis/passes/bools"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/analysis/passes/buildtag"
	"golang.org/x/tools/go/analysis/passes/cgocall"
	"golang.org/x/tools/go/analysis/passes/composite"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"golang.org/x/tools/go/analysis/passes/ctrlflow"
	"golang.org/x/tools/go/analysis/passes/deepequalerrors"
	"golang.org/x/tools/go/analysis/passes/directive"
	"golang.org/x/tools/go/analysis/passes/errorsas"
	"golang.org/x/tools/go/analysis/passes/fieldalignment"
	"golang.org/x/tools/go/analysis/passes/findcall"
	"golang.org/x/tools/go/analysis/passes/framepointer"
	"golang.org/x/tools/go/analysis/passes/httpresponse"
	"golang.org/x/tools/go/analysis/passes/ifaceassert"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/passes/loopclosure"
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"golang.org/x/tools/go/analysis/passes/nilness"
	"golang.org/x/tools/go/analysis/passes/pkgfact"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/reflectvaluecompare"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/shift"
	"golang.org/x/tools/go/analysis/passes/sigchanyzer"
	"golang.org/x/tools/go/analysis/passes/sortslice"
	"golang.org/x/tools/go/analysis/passes/stdmethods"
	"golang.org/x/tools/go/analysis/passes/stringintconv"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"golang.org/x/tools/go/analysis/passes/testinggoroutine"
	"golang.org/x/tools/go/analysis/passes/tests"
	"golang.org/x/tools/go/analysis/passes/timeformat"
	"golang.org/x/tools/go/analysis/passes/unmarshal"
	"golang.org/x/tools/go/analysis/passes/unreachable"
	"golang.org/x/tools/go/analysis/passes/unsafeptr"
	"golang.org/x/tools/go/analysis/passes/unusedresult"
	"golang.org/x/tools/go/analysis/passes/unusedwrite"
	"golang.org/x/tools/go/analysis/passes/usesgenerics"
)

func main() {
	multichecker.Main(

		asmdecl.Analyzer, // checks that assembly functions are declared correctly
		assign.Analyzer,  // checks for useless assignments
		atomic.Analyzer,  // checks for common mistaken usages of the sync/atomic package
		atomicalign.Analyzer,
		bools.Analyzer,     // checks for boolean expressions that can be simplified
		buildssa.Analyzer,  // builds the SSA representation of a package
		buildtag.Analyzer,  // checks that +build tags are valid and correctly positioned
		cgocall.Analyzer,   // checks that cgo calls are valid
		composite.Analyzer, // checks that composite literals are valid
		copylock.Analyzer,  // checks for locks inadvertently passed by value
		ctrlflow.Analyzer,  // checks for unreachable code
		deepequalerrors.Analyzer,
		directive.Analyzer,
		errorsas.Analyzer, // checks that errors are not compared with == or != (use errors.Is or errors.As instead)
		fieldalignment.Analyzer,
		findcall.Analyzer,     // checks that calls to runtime.SetFinalizer are valid
		framepointer.Analyzer, // checks that the framepointer buildmode is used when using cgo
		httpresponse.Analyzer, // checks that http response bodies are closed
		ifaceassert.Analyzer,  // checks that type assertions on interfaces are valid
		inspect.Analyzer,      // runs an arbitrary Go analysis function
		loopclosure.Analyzer,  // checks that loop closures do not capture loop variables
		lostcancel.Analyzer,   // checks that context cancelation signals are propagated
		nilfunc.Analyzer,      // checks that nil function values are not called
		nilness.Analyzer,
		pkgfact.Analyzer, // checks that package facts are valid
		printf.Analyzer,  // checks that printf-like invocations are valid
		reflectvaluecompare.Analyzer,
		shadow.Analyzer,
		shift.Analyzer,       // checks that shifts are valid
		sigchanyzer.Analyzer, // checks that signal handlers are valid
		sortslice.Analyzer,
		stdmethods.Analyzer,    // checks that methods on the standard library types are valid
		stringintconv.Analyzer, // checks that string conversions are valid
		structtag.Analyzer,     // checks that struct tags are valid
		testinggoroutine.Analyzer,
		tests.Analyzer, // checks that test files are valid
		timeformat.Analyzer,
		unmarshal.Analyzer,    // checks that unmarshal calls are valid
		unreachable.Analyzer,  // checks that unreachable code is not executed
		unsafeptr.Analyzer,    // checks that unsafe pointer conversions are valid
		unusedresult.Analyzer, // checks that results of calls to certain functions are used
		unusedwrite.Analyzer,
		usesgenerics.Analyzer,
	)
}
