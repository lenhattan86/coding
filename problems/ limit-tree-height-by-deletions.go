// Given an N-ary tree where each node has an integer id and a list of children, calculate the maximum height of the tree after removing certain nodes. When a node is deleted, its children are promoted and connected directly to the deleted node's parent, similar to removing a node from a linked list while maintaining the chain. The tree height is defined as the number of nodes on the longest path from the root to any leaf node. The root node will never be deleted, and a single-node tree has height 1.

package main

