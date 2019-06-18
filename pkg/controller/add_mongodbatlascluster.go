package controller

import (
	"github.com/Knappek/mongodbatlas-operator/pkg/controller/mongodbatlascluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, mongodbatlascluster.Add)
}
