package main

import "fmt"

type Config struct {
	SelectedClusterID int       `json:"selectedClusterID"`
	Clusters          []Cluster `json:"clusters"`
}
type Cluster struct {
	ID             int    `json:"clusterId"`
	Name           string `json:"clusterName"`
	SelectedNodeID int    `json:"selectedNodeID"`
	Nodes          []Node `json:"nodes"`
}
type Node struct {
	NodeID    int    `json:"nodeId"`
	Name      string `json:"name"`
	IP        string `json:"ip"`
	LoginType int    `json:"loginType"`
	Password  string `json:"password,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}
type tst struct {
	Str Config
}

func main() {

	data := Config{}
	err := readConfigJSON(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
