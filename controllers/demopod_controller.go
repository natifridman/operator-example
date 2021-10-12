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

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	examplev1 "operator-example/api/v1"
)

// DemoPodReconciler reconciles a DemoPod object
type DemoPodReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=example.nati.com,resources=demopods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=example.nati.com,resources=demopods/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=example.nati.com,resources=demopods/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DemoPod object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile
func (r *DemoPodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	l.Info("Enter Reconcile", "req", req)

	pod := &examplev1.DemoPod{}
	r.Get(ctx, types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, pod)

	l.Info("Enter Reconcile", "spec", pod.Spec, "status", pod.Status)

	if pod.Spec.Name != pod.Status.Name {
		pod.Status.Name = pod.Spec.Name
		r.Status().Update(ctx, pod)
	}

	podFound := &v1.Pod{}
	err := r.Get(ctx, types.NamespacedName{Name: pod.Spec.Name, Namespace: req.Namespace}, podFound)

	if err != nil && errors.IsNotFound(err) {
		newPod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: pod.Namespace,
				Name:      pod.Spec.Name,
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{{
					Image: pod.Spec.Image,
					Name:  pod.Spec.Name,
				}},
			},
		}
		l.Info("Creating new Pod", "pod", newPod)

		err = r.Create(ctx, newPod)
		if err != nil {
			l.Error(err, "Failed to create new Pod", "Namespace", req.Namespace, "Name", pod.Spec.Name)
		}

		l.Info("Pod created", "pod", newPod)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DemoPodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&examplev1.DemoPod{}).
		Complete(r)
}
