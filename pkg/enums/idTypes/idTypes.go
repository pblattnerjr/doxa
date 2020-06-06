//go:generate enumer -type=IDType -json -text -yaml -sql
// 1) go get github.com/alvaroloes/enumer
// 2) in the enum subfolder for this enum: go generate

// Package modes provides an enum of ID types.
// IDType indicates NID (no ID, use literal), RID (relative ID), or SID (Specific ID)
package idTypes

type IDType int
const (
	NA IDType = iota
	NID
	RID
	SID
)


