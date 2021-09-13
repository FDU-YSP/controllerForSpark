/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	_ "sigs.k8s.io/controller-runtime/pkg/event"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	sparkv1 "my.domain/controllerForSpark/api/v1"
)

// SparkJobReconciler reconciles a SparkJob object
type SparkJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

const (
	SparkStatusPendingPhase  = "pending"
	SparkStatusCreatingPhase = "creating"
	SparkStatusRunningPhase  = "running"
	SparkStatusSucceedPhase  = "succeed"
	SparkStatusFailedPhase   = "failed"
	SparkStatusUnknowPhase   = "unknow"
)

//+kubebuilder:rbac:groups=spark.my.domain,resources=sparkjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=spark.my.domain,resources=sparkjobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=spark.my.domain,resources=sparkjobs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SparkJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *SparkJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	//ctx2 := context.Background()
	// your logic here
	var  sparkJob sparkv1.SparkJob

	// req.NamespacedName = namespace + / + name
	if err := r.Get(ctx, req.NamespacedName, &sparkJob); err != nil {
		// r.Log.WithName("unable to fetch migrate")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	statusOfCopy := sparkJob.Status.DeepCopy()
	// specOfCopy   := sparkJob.Spec.DeepCopy()

	defer func() {
		// print log

		// change current object's status
		sparkJob.Status = *statusOfCopy
		if err := r.Status().Update(ctx, &sparkJob); err != nil {
			fmt.Println(err)
		}

		if err := r.Update(ctx, &sparkJob); err != nil {
			fmt.Println(err)
		}

	}()

	switch statusOfCopy.Phase {
	case "":
		// print log

		// your logic here
		statusOfCopy.Phase = SparkStatusPendingPhase
		fmt.Println("pending")

		return ctrl.Result{Requeue: true}, nil
	case SparkStatusPendingPhase:
		// print log
		// your logic here
		statusOfCopy.Phase = SparkStatusCreatingPhase
		fmt.Println("creating")

		return ctrl.Result{Requeue: true}, nil

	case SparkStatusCreatingPhase:
		statusOfCopy.Phase = SparkStatusRunningPhase
		fmt.Println("running")

		return ctrl.Result{Requeue: true}, nil
	case SparkStatusRunningPhase:
		// print log
		// your logic here
		statusOfCopy.Phase = SparkStatusSucceedPhase
		fmt.Println("succeed")

		return ctrl.Result{Requeue: true}, nil
	}
	return ctrl.Result{Requeue: true}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SparkJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sparkv1.SparkJob{}).
		Complete(r)
}

//
//type SparkJobPredicate struct {
//}
//
//func (s SparkJobPredicate) Update(e event.UpdateEvent) bool {
//	return false
//}
//func (s SparkJobPredicate) Delete(e event.DeleteEvent) bool {
//	return false
//}
//func (s SparkJobPredicate) Generic(e event.GenericEvent) bool {
//	return false
//}