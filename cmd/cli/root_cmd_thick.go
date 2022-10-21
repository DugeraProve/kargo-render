//go:build thick
// +build thick

package main

import "fmt"

func init() {
	desc = fmt.Sprintf(
		"%s.\n\nThis \"thick\" variant of the Bookkeeper CLI does not offload "+
			"work to a Bookkeeper server.\n\nTHIS CLI IS UNLIKELY TO WORK OUTSIDE "+
			"THE CONTEXT OF THE OFFICIAL BOOKKEEPER DOCKER IMAGE.",
		desc,
	)
}
