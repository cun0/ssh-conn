package main

import "io"

type SwitchNodeOp struct {
	NodeName string
	NodeID   int
}
type SwitchClusterOp struct {
	Target string
}

func (op SwitchNodeOp) Run(_, stderr io.Writer) error {
	//var newNode string
	var err error

	return err
}

func switchNode(nodeName string, nodeID int) {

}
