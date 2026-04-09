package storage

import "github.com/papercomputeco/tapes/pkg/merkle"

// Chain is the result of walking a node's parent edges back toward a root.
//
// A Chain whose Incomplete field is false reached a legitimate root (a node
// whose parent_hash is nil or empty). A Chain whose Incomplete field is true
// stopped because a parent_hash pointed at a node that is not currently
// present in this store — see MissingParent. The nodes in Nodes are still
// valid; the store simply can't resolve any higher ancestors from here.
//
// This state is expected on large or long-lived stores that trim/offload
// older data, merge content from foreign sources, or receive chains whose
// ancestors live in another store altogether. Callers should treat it as
// informational (a signal to render a marker, or trigger a future "thaw"
// lookup against another source), not as corruption.
type Chain struct {
	// Nodes is the walk output in node-first order: index 0 is the node
	// the walk started from, and the last element is either a real root
	// or the last resolvable node whose parent could not be found.
	Nodes []*merkle.Node

	// Incomplete is true when the walk stopped at a missing parent.
	Incomplete bool

	// MissingParent is the parent_hash that could not be resolved in this
	// store. Only set when Incomplete is true.
	MissingParent string
}

// Complete reports whether the walk reached a real root.
func (c *Chain) Complete() bool {
	return c != nil && !c.Incomplete
}
