package cmd

import (
	"flag"
	"fmt"
	"strings"

	"github.com/VKCOM/noverify/src/lintdoc"
	"github.com/VKCOM/noverify/src/linter"
	"github.com/VKCOM/noverify/src/rules"
)

func Help(*MainConfig) (int, error) {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Printf("Usage of noverify:\n")
		fmt.Printf("  $ noverify [command] -stubs-dir=/path/to/phpstorm-stubs -cache-dir=/cache/dir /project/root\n\n")
		GlobalCmds.PrintHelpPage()
		return 0, nil
	}

	mainSubject := args[0]
	switch mainSubject {
	case "checkers":
		declareRules()
		var subSubject string
		if len(args) > 1 {
			subSubject = args[1]
		}
		if subSubject == "" {
			helpAllCheckers()
			return 0, nil
		}
		err := helpChecker(subSubject)
		if err != nil {
			return 1, err
		}
		return 0, err
	default:
		return 1, fmt.Errorf("unknown subject: %s", mainSubject)
	}
}

func declareRules() {
	p := rules.NewParser()
	linter.Rules = rules.NewSet()
	ruleSets, err := InitEmbeddedRules(p, func(r rules.Rule) bool { return true })
	if err != nil {
		panic(err)
	}
	for _, rset := range ruleSets {
		linter.DeclareRules(rset)
	}
}

func helpAllCheckers() {
	for _, info := range linter.GetDeclaredChecks() {
		fmt.Printf("%s: %s\n", info.Name, info.Comment)
	}
}

func helpChecker(checkerName string) error {
	var info linter.CheckInfo
	checks := linter.GetDeclaredChecks()
	for i := range checks {
		if checks[i].Name == checkerName {
			info = checks[i]
		}
	}
	if info.Name == "" {
		return fmt.Errorf("checker %s not found", checkerName)
	}
	var buf strings.Builder
	if err := lintdoc.RenderCheckDocumentation(&buf, info); err != nil {
		return err
	}
	fmt.Println(buf.String())
	return nil
}
