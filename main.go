package main

import (
	"log"
	"os"

	"github.com/goreleaser/nfpm/v2"

	// Register the rpm packager with the nfpm framework.
	_ "github.com/goreleaser/nfpm/v2/rpm"
)

func main() {
	info := nfpm.WithDefaults(
		&nfpm.Info{
			Name:    "meta",
			Version: "1.0",
		},
	)
	// For RPM, see https://rpm-software-management.github.io/rpm/manual/dependencies.html#requires
	info.Overridables.Depends = []string{
		"foo >= 1.0",
		"foo <  2.0",
	}
	// For RPM, see https://rpm-software-management.github.io/rpm/manual/dependencies.html#conflicts
	info.Overridables.Conflicts = []string{
		"bar",
	}

	rpmPackager, err := nfpm.Get("rpm")
	if err != nil {
		log.Fatal(err)
	}

	err = rpmPackager.Package(info, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
