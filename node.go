package mme

type Node struct {
	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node
}
