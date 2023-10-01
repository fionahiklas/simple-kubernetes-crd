package controllers

import (
	"context"
	"github.com/fionahiklas/simple-kubernetes-crd/etoperator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// ensureService ensures Service is Running in a namespace.
func (etr *ElectricTreesReconciler) ensureService(request reconcile.Request,
	treeInstance *v1alpha1.ElectricTrees,
	serviceToCreate *corev1.Service) (*reconcile.Result, error) {

	// See if service already exists and create if it doesn't
	found := &appsv1.Deployment{}

	err := etr.Get(context.TODO(), types.NamespacedName{
		Name:      serviceToCreate.Name,
		Namespace: treeInstance.Namespace,
	}, found)

	if err != nil && errors.IsNotFound(err) {

		// Create the service
		err = etr.Create(context.TODO(), serviceToCreate)

		if err != nil {
			// Service creation failed, requeue
			return &reconcile.Result{}, err

		} else {
			// Service creation was successful
			return nil, nil
		}

	} else if err != nil {
		// Error that isn't due to the service not existing, requeue
		return &reconcile.Result{}, err
	}

	return nil, nil
}

func (etr *ElectricTreesReconciler) electricTreeService(tree *v1alpha1.ElectricTrees) *corev1.Service {
	labels := labels(tree)

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tree.Name + "-svc",
			Namespace: tree.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{{
				Protocol:   corev1.ProtocolTCP,
				Port:       7777,
				TargetPort: intstr.FromString("tree-port"),
				NodePort:   32707,
			}},
			Type: corev1.ServiceTypeNodePort,
		},
	}

	controllerutil.SetControllerReference(tree, service, etr.Scheme)
	return service
}
