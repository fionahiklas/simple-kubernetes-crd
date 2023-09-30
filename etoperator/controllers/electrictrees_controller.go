/*
   Copyright 2023 Fiona Bianchi.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package controllers

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	hiklascomv1alpha1 "github.com/fionahiklas/simple-kubernetes-crd/etoperator/api/v1alpha1"
)

// ElectricTreesReconciler reconciles a ElectricTrees object
type ElectricTreesReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=hiklas.com,resources=electrictrees,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hiklas.com,resources=electrictrees/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hiklas.com,resources=electrictrees/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (etr *ElectricTreesReconciler) Reconcile(ctx context.Context, ctrlRequest ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx).WithValues("ElectricTrees", ctrlRequest.NamespacedName)

	// Fetch the Traveller instance
	instance := &hiklascomv1alpha1.ElectricTrees{}
	err := etr.Get(context.TODO(), ctrlRequest.NamespacedName, instance)

	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. Return and don't requeue
			return reconcile.Result{}, nil
		}

		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Check if this Deployment already exists
	found := &appsv1.Deployment{}
	err = etr.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, found)

	var result *reconcile.Result
	result, err = etr.ensureDeployment(ctrlRequest, instance, etr.electricTreeDeployment(instance))

	if result != nil {
		log.Error(err, "Deployment Not ready")
		return *result, err
	}

	// Check if this Service already exists
	result, err = etr.ensureService(ctrlRequest, instance, etr.electricTreeService(instance))
	if result != nil {
		log.Error(err, "Service Not ready")
		return *result, err
	}

	// Deployment and Service already exists - don't requeue
	log.Info("Skip reconcile: Deployment and service already exists",
		"Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (etr *ElectricTreesReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hiklascomv1alpha1.ElectricTrees{}).
		Complete(etr)
}
