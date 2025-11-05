package handlers

import (
	"context"
	"fmt"
	"gok8s/kubernetsServ"
	"io"
	"log"
	"net/http"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
)

func DeployNamespaceHandler(w http.ResponseWriter, r *http.Request) {

	// 1. Read request data
	yamlData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failure to get request data:  %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// check YAML content is empty or not.
	if len(yamlData) == 0 {
		http.Error(w, "YAML content can not be empty", http.StatusBadRequest)
		return
	}

	// 2. Process data
	log.Printf("Try to create dynamic client: %v", kubernetsServ.ConfigR)
	// decode YAML to Unstructured
	//yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	decode := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

	obj := &unstructured.Unstructured{}
	_, _, err = decode.Decode(yamlData, nil, obj)
	if err != nil {
		log.Printf("Decode YAML file failure:  %v", err)
		http.Error(w, fmt.Sprintf("Decode yaml file failure: %v", err), http.StatusBadRequest)
		return
	}
	log.Printf("Sucess decode yaml file, the resource type is: %s", obj.GetKind())

	// Verify that the YAML file contains a "kind" parameter.
	if obj.GetKind() == "" {
		http.Error(w, "YAML file lack kind parameter.", http.StatusBadRequest)
		return
	}

	// Get the resource type from the YAML file.
	kind := obj.GetKind()
	resource := strings.ToLower(kind)
	if !strings.HasSuffix(resource, "s") {
		resource += "s" //  Deployment -> deployments
	}

	// Get namespace and default is default
	namespace := obj.GetNamespace()
	if namespace == "" {
		namespace = "default"
	} else {
		namespace = strings.ToLower(namespace)
	}

	namespace1 := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: obj.GetName(),
			Labels: map[string]string{
				"app": func() string {
					if obj.GetNamespace() == "" {
						return "app"
					}
					return obj.GetLabels()["app"]
				}(),
				"environment": obj.GetLabels()["environment"],
			},

			Annotations: map[string]string{
				"description": obj.GetAnnotations()["description"],
				"managed-by":  obj.GetAnnotations()["managed-by"],
			},
		},
	}
	_, err = kubernetsServ.Clientset.CoreV1().Namespaces().Create(context.TODO(), namespace1, metav1.CreateOptions{})
	if r.Method != http.MethodPost {
		http.Error(w, "Can not support this type request.", http.StatusMethodNotAllowed)
		return
	}

}
