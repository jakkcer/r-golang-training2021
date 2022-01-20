package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func packages(patterns []string) []string {
	args := []string{"list", "-f={{.ImportPath}}"}
	for _, pkg := range patterns {
		args = append(args, pkg)
	}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return strings.Fields(string(out))
}

func ancestors(packageNames []string) []string {
	targets := make(map[string]bool)
	for _, pkg := range packageNames {
		targets[pkg] = true
	}

	args := []string{"list", `-f={{.ImportPath}} {{join .Deps " "}}`, "..."}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	var pkgs []string
	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		fields := strings.Fields(s.Text())
		pkg := fields[0]
		deps := fields[1:]
		for _, dep := range deps {
			if targets[dep] {
				pkgs = append(pkgs, pkg)
				break
			}
		}
	}
	return pkgs
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	pkgs := ancestors(packages(os.Args[1:]))
	sort.StringSlice(pkgs).Sort()
	for _, pkg := range pkgs {
		fmt.Println(pkg)
	}
}
