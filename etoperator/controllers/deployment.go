package controllers

import (
	"github.com/fionahiklas/simple-kubernetes-crd/etoperator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (etr *ElectricTreesReconciler) electricTreeDeployment(tree *v1alpha1.ElectricTrees) *appsv1.Deployment {
	labels := labels(tree)
	size := int32(1)

	deploymentObject := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tree.Name,
			Namespace: tree.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &size,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:           "electrictree:0.0.1",
						ImagePullPolicy: corev1.PullAlways,
						Name:            "electric-tree-pod",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 8080,
							Name:          "tree",
						}},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(tree, deploymentObject, etr.Scheme)
	return deploymentObject
}
