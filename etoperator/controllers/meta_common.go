package controllers

import "github.com/fionahiklas/simple-kubernetes-crd/etoperator/api/v1alpha1"

func labels(tree *v1alpha1.ElectricTrees) map[string]string {
	return map[string]string{
		"app":       "electric_trees_app",
		"tree_name": tree.Name,
	}
}
