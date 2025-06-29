package main

import (
	"fmt"
	"testing"

	"github.com/nmaupu/depth"
)

func Test_parse(t *testing.T) {
	tests := []struct {
		internal bool
		test     bool
		depth    int
		json     bool
		explain  string
	}{
		{true, true, 0, true, ""},
		{false, false, 10, false, ""},
		{true, false, 10, false, ""},
		{false, true, 5, true, ""},
		{false, true, 5, true, "github.com/nmaupu/depth"},
		{false, true, 5, true, ""},
	}

	for idx, tt := range tests {
		tr, _ := parse([]string{
			fmt.Sprintf("-internal=%v", tt.internal),
			fmt.Sprintf("-test=%v", tt.test),
			fmt.Sprintf("-max=%v", tt.depth),
			fmt.Sprintf("-json=%v", tt.json),
			fmt.Sprintf("-explain=%v", tt.explain),
		})

		if tr.ResolveInternal != tt.internal {
			t.Fatalf("[%v] Unexpected ResolveInternal, expected=%v, got=%v", idx, tt.internal, tr.ResolveInternal)
		} else if tr.ResolveTest != tt.test {
			t.Fatalf("[%v] Unexpected ResolveTest, expected=%v, got=%v", idx, tt.test, tr.ResolveTest)
		} else if tr.MaxDepth != tt.depth {
			t.Fatalf("[%v] Unexpected MaxDepth, expected=%v, got=%v", idx, tt.depth, tr.MaxDepth)
		} else if outputJSON != tt.json {
			t.Fatalf("[%v] Unexpected outputJSON, expected=%v, got=%v", idx, tt.json, outputJSON)
		} else if explainPkg != tt.explain {
			t.Fatalf("[%v] Unexpected explainPkg, expected=%v, got=%v", idx, tt.explain, explainPkg)
		}
	}
}

func Example_handlePkgsStrings() {
	var t depth.Tree

	handlePkgs(&t, []string{"strings"}, false, "")
	// Output:
	// strings
	//   ├ errors
	//   ├ internal/bytealg
	//   ├ io
	//   ├ sync
	//   ├ unicode
	//   ├ unicode/utf8
	//   └ unsafe
	// 7 dependencies (7 internal, 0 external, 0 testing).
}

func Example_handlePkgsTestStrings() {
	var t depth.Tree
	t.ResolveTest = true

	handlePkgs(&t, []string{"strings"}, false, "")
	// Output:
	// strings
	//   ├ bytes
	//   ├ errors
	//   ├ fmt
	//   ├ internal/bytealg
	//   ├ internal/testenv
	//   ├ io
	//   ├ math/rand
	//   ├ reflect
	//   ├ strconv
	//   ├ sync
	//   ├ testing
	//   ├ unicode
	//   ├ unicode/utf8
	//   └ unsafe
	// 14 dependencies (14 internal, 0 external, 7 testing).
}

func Example_handlePkgsDepth() {
	var t depth.Tree

	handlePkgs(&t, []string{"github.com/nmaupu/depth/cmd/depth"}, false, "")
	// Output:
	// github.com/nmaupu/depth/cmd/depth
	//   ├ encoding/json
	//   ├ flag
	//   ├ fmt
	//   ├ io
	//   ├ os
	//   ├ strings
	//   └ github.com/nmaupu/depth
	//     ├ bytes
	//     ├ errors
	//     ├ go/build
	//     ├ os
	//     ├ path
	//     ├ sort
	//     └ strings
	// 12 dependencies (11 internal, 1 external, 0 testing).
}

func Example_handlePkgsUnknown() {
	var t depth.Tree

	handlePkgs(&t, []string{"notreal"}, false, "")
	// Output:
	// 'notreal': FATAL: unable to resolve root package
}

func Example_handlePkgsJson() {
	var t depth.Tree
	handlePkgs(&t, []string{"strings"}, true, "")

	// Output:
	// {
	//   "name": "strings",
	//   "internal": true,
	//   "resolved": true,
	//   "deps": [
	//     {
	//       "name": "errors",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "internal/bytealg",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "io",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "sync",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "unicode",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "unicode/utf8",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "unsafe",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     }
	//   ]
	// }

}

func Example_handlePkgsExplain() {
	var t depth.Tree

	handlePkgs(&t, []string{"github.com/nmaupu/depth/cmd/depth"}, false, "strings")
	// Output:
	// github.com/nmaupu/depth/cmd/depth -> strings
	// github.com/nmaupu/depth/cmd/depth -> github.com/nmaupu/depth -> strings
}
