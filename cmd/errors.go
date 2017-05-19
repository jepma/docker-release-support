/*
Package gitwrap implements a simple library for GIT interactions.
*/
package cmd

import "errors"

// Error codes returned by failures to parse an expression.
var (
	ErrChangesPending = errors.New("repository contains changes, cannot release patch with a dirty tree")
	ErrNoChanges      = errors.New("repository contains no changes, why would you want to tag?")
)
