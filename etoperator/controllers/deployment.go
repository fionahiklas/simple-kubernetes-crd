package controllers

import (
	"context"
	"github.com/fionahiklas/simple-kubernetes-crd/etoperator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func (etr *ElectricTreesReconciler) ensureDeployment(request reconcile.Request,
	treeInstance *v1alpha1.ElectricTrees,
	deploymentToCreate *appsv1.Deployment) (*reconcile.Result, error) {

	// See if deployment already exists and create if it doesn't
	existingDeployment := &appsv1.Deployment{}

	err := etr.Get(context.TODO(), types.NamespacedName{
		Name:      deploymentToCreate.Name,
		Namespace: treeInstance.Namespace,
	}, existingDeployment)

	if err != nil && errors.IsNotFound(err) {

		// Create the deployment
		err = etr.Create(context.TODO(), deploymentToCreate)

		if err != nil {
			// Deployment failed, requeue to try again
			return &reconcile.Result{}, err

		} else {
			// Deployment was successful
			return nil, nil
		}

	} else if err != nil {
		// Error that isn't due to the deployment not existing, requeue
		return &reconcile.Result{}, err
	}

	return nil, nil
}

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
						Image:           "electrictrees:latest",
						ImagePullPolicy: corev1.PullAlways,
						Name:            "electric-trees-pod",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 7777,
							Name:          "treePort",
						}},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(tree, deploymentObject, etr.Scheme)
	return deploymentObject
}
