package controllers

import (
	"github.com/fionahiklas/simple-kubernetes-crd/etoperator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (etr *ElectricTreesReconciler) backendService(tree *v1alpha1.ElectricTrees) *corev1.Service {
	labels := labels(tree)

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tree.Name,
			Namespace: tree.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{{
				Protocol:   corev1.ProtocolTCP,
				Port:       80,
				TargetPort: intstr.FromInt(8080),
				NodePort:   30685,
			}},
			Type: corev1.ServiceTypeNodePort,
		},
	}

	controllerutil.SetControllerReference(tree, service, etr.Scheme)
	return service
}
