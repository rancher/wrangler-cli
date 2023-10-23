package types

import (
	"fmt"
	"strings"

	"github.com/rancher/wrangler/v2/pkg/gvk"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	meta1 "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	DefaultVersion = "v0"

	ConfigType          = "configmap"
	ServiceType         = "service"
	PodType             = "pod"
	DeploymentType      = "deploy"
	DaemonSetType       = "ds"
	NamespaceType       = "namespace"
	RouterType          = "router"
	ExternalServiceType = "externalservice"
	SecretType          = "secret"
	BuildType           = "taskrun"
	StackType           = "stack"
)

var (
	Aliases = map[string]string{
		"config":           ConfigType,
		"configs":          ConfigType,
		"configmaps":       ConfigType,
		"svc":              ServiceType,
		"svcs":             ServiceType,
		"pods":             PodType,
		"deployment":       DeploymentType,
		"deployments":      DeploymentType,
		"deploys":          DeploymentType,
		"routers":          RouterType,
		"externalservices": ExternalServiceType,
		"secrets":          SecretType,
		"build":            BuildType,
		"taskruns":         BuildType,
		"stacks":           StackType,
		"ns":               NamespaceType,
		"namespace":        NamespaceType,
	}
)

type Resource struct {
	LookupName, Name, Namespace, Type string
	App, Version                      string
	Object                            runtime.Object
}

func (r Resource) String() string {
	return r.StringDefaultNamespace("")
}

func (r Resource) StringDefaultNamespace(defaultNamespace string) string {
	if r.LookupName != "" {
		return r.LookupName
	}
	buf := strings.Builder{}
	if defaultNamespace == "" || (r.Namespace != "" && r.Namespace != defaultNamespace) {
		buf.WriteString(r.Namespace)
		buf.WriteString(":")
	}

	if r.Type != ServiceType {
		buf.WriteString(r.Type)
		buf.WriteString("/")
	}

	if r.Type == ServiceType {
		buf.WriteString(r.App)
		if r.Version != DefaultVersion {
			buf.WriteString("@")
			buf.WriteString(r.Version)
		}
	} else {
		buf.WriteString(r.Name)
	}

	return buf.String()
}

func FromObject(obj runtime.Object) (Resource, error) {
	result := Resource{}
	switch obj.(type) {
	case *corev1.Secret:
		result.Type = SecretType
	case *corev1.Pod:
		result.Type = PodType
	case *appsv1.Deployment:
		result.Type = DeploymentType
	case *appsv1.DaemonSet:
		result.Type = DaemonSetType
	case *corev1.ConfigMap:
		result.Type = ConfigType
	default:
		gvk, err := gvk.Get(obj)
		if err != nil {
			return Resource{}, fmt.Errorf("unrecognized type: %w", err)
		}
		result.Type = strings.ToLower(gvk.Kind)
	}

	meta, err := meta1.Accessor(obj)
	if err != nil {
		return result, err
	}

	result.Namespace = meta.GetNamespace()
	result.Name = meta.GetName()

	return result, nil
}
