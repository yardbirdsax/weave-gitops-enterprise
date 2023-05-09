package configuration

import (
	"fmt"
	"github.com/fluxcd/helm-controller/api/v2beta1"
	"github.com/fluxcd/kustomize-controller/api/v1beta2"
	sourcev1 "github.com/fluxcd/source-controller/api/v1beta2"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ObjectKind struct {
	Gvk                 schema.GroupVersionKind     `json:"groupVersionKind"`
	NewClientObjectFunc func() client.Object        `json:"-"`
	AddToSchemeFunc     func(*runtime.Scheme) error `json:"-"`
}

func (ok ObjectKind) String() string {
	return ok.Gvk.String()
}

func (o ObjectKind) Validate() error {
	if o.Gvk.Kind == "" {
		return fmt.Errorf("missing gvk")
	}
	if o.NewClientObjectFunc == nil {
		return fmt.Errorf("missing client func")
	}
	if o.AddToSchemeFunc == nil {
		return fmt.Errorf("missing add to scheme func")
	}

	return nil
}

var (
	HelmReleaseObjectKind = ObjectKind{
		Gvk: v2beta1.GroupVersion.WithKind(v2beta1.HelmReleaseKind),
		NewClientObjectFunc: func() client.Object {
			return &v2beta1.HelmRelease{}
		},
		AddToSchemeFunc: v2beta1.AddToScheme,
	}
	KustomizationObjectKind = ObjectKind{
		Gvk: v1beta2.GroupVersion.WithKind(v1beta2.KustomizationKind),
		NewClientObjectFunc: func() client.Object {
			return &v1beta2.Kustomization{}
		},
		AddToSchemeFunc: v1beta2.AddToScheme,
	}
	HelmRepositoryObjectKind = ObjectKind{
		Gvk: sourcev1.GroupVersion.WithKind(sourcev1.HelmRepositoryKind),
		NewClientObjectFunc: func() client.Object {
			return &sourcev1.HelmRepository{}
		},
		AddToSchemeFunc: sourcev1.AddToScheme,
	}
	HelmChartObjectKind = ObjectKind{
		Gvk: sourcev1.GroupVersion.WithKind(sourcev1.HelmChartKind),
		NewClientObjectFunc: func() client.Object {
			return &sourcev1.HelmChart{}
		},
		AddToSchemeFunc: sourcev1.AddToScheme,
	}
	GitRepositoryObjectKind = ObjectKind{
		Gvk: sourcev1.GroupVersion.WithKind(sourcev1.GitRepositoryKind),
		NewClientObjectFunc: func() client.Object {
			return &sourcev1.GitRepository{}
		},
		AddToSchemeFunc: sourcev1.AddToScheme,
	}
	OCIRepositoryObjectKind = ObjectKind{
		Gvk: sourcev1.GroupVersion.WithKind(sourcev1.OCIRepositoryKind),
		NewClientObjectFunc: func() client.Object {
			return &sourcev1.OCIRepository{}
		},
		AddToSchemeFunc: sourcev1.AddToScheme,
	}
	BucketObjectKind = ObjectKind{
		Gvk: sourcev1.GroupVersion.WithKind(sourcev1.BucketKind),
		NewClientObjectFunc: func() client.Object {
			return &sourcev1.Bucket{}
		},
		AddToSchemeFunc: sourcev1.AddToScheme,
	}
	RoleObjectKind = ObjectKind{
		Gvk: rbacv1.SchemeGroupVersion.WithKind("Role"),
		NewClientObjectFunc: func() client.Object {
			return &rbacv1.Role{}
		},
		AddToSchemeFunc: rbacv1.AddToScheme,
	}
	ClusterRoleObjectKind = ObjectKind{
		Gvk: rbacv1.SchemeGroupVersion.WithKind("ClusterRole"),
		NewClientObjectFunc: func() client.Object {
			return &rbacv1.ClusterRole{}
		},
		AddToSchemeFunc: rbacv1.AddToScheme,
	}
	RoleBindingObjectKind = ObjectKind{
		Gvk: rbacv1.SchemeGroupVersion.WithKind("RoleBinding"),
		NewClientObjectFunc: func() client.Object {
			return &rbacv1.RoleBinding{}
		},
		AddToSchemeFunc: rbacv1.AddToScheme,
	}
	ClusterRoleBindingObjectKind = ObjectKind{
		Gvk: rbacv1.SchemeGroupVersion.WithKind("ClusterRoleBinding"),
		NewClientObjectFunc: func() client.Object {
			return &rbacv1.ClusterRoleBinding{}
		},
		AddToSchemeFunc: rbacv1.AddToScheme,
	}
)

// SupportedObjectKinds list with the default supported Object resources to query.
var SupportedObjectKinds = []ObjectKind{
	HelmReleaseObjectKind,
	KustomizationObjectKind,
	HelmRepositoryObjectKind,
	HelmChartObjectKind,
	GitRepositoryObjectKind,
	OCIRepositoryObjectKind,
	BucketObjectKind,
}

// SupportedRbacKinds list with the default supported RBAC resources.
var SupportedRbacKinds = []ObjectKind{
	RoleObjectKind,
	ClusterRoleObjectKind,
	RoleBindingObjectKind,
	ClusterRoleBindingObjectKind,
}
