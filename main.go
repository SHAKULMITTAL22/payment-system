package main

import (
	"fmt"
	"log"

	"golang.org/x/tools/go/packages"
)

func main() {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedImports | packages.NeedSyntax,
	}

	pkgs, err := packages.Load(cfg, "github.com/SHAKULMITTAL22/payment-system/payment")
	if err != nil {
		log.Fatalf("Failed to load package with error: %v", err)
	}
	for _, p := range pkgs {
		fmt.Printf("Package: %s\n", p.PkgPath)
		fmt.Println("Types in scope:")
		for _, name := range p.Types.Scope().Names() {
			fmt.Println(name)
			obj := p.Types.Scope().Lookup(name)
			if obj != nil && obj.Name() == "Payment" {
				fmt.Println("Found Payment struct.")
				break
			}
		}
	}
}
