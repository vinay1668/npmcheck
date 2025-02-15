package main

import (
	"fmt"
	"npmcheck/pkg/npmcheck"
)

func main() {
	results, err := npmcheck.CheckPackageVersions("package.json")
	if err != nil {
		fmt.Printf("Error checking versions: %v\n", err)
		return
	}

	for pkg, versions := range results {
		fmt.Printf("Package: %s\n", pkg)
		fmt.Printf("Current Version: %s\n", versions.Current)
		fmt.Printf("Latest Version: %s\n", versions.Latest)
		fmt.Println("-------------------")
	}
}
