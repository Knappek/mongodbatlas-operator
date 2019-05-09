package controller

import (
	"github.com/Knappek/mongodbatlas-operator/pkg/controller/mongodbatlasproject"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, mongodbatlasproject.Add)
}
