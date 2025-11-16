package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"gok8s/kubernetsServ"
	"io"
	"log"
	"net/http"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/dynamic"
)

func DeployPodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Can not support this type request.", http.StatusMethodNotAllowed)
		return
	}

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
	dynamicClient, err := dynamic.NewForConfig(kubernetsServ.ConfigR)
	if err != nil {
		log.Printf("Create dynamic resource failure:  %v", err)
		http.Error(w, fmt.Sprintf("Create Kubernetes client %v", err), http.StatusInternalServerError)
		return
	}

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

	// Creat a dynamic client for the resource.
	gvr := obj.GroupVersionKind()
	log.Printf("Ready to create resource, type: %s, nameSpace is: %s", resource, namespace)
	resourceClient := dynamicClient.Resource(gvr.GroupVersion().WithResource(resource)).Namespace(namespace)

	// Create the resource.
	log.Printf("Sending create request...")
	result, err := resourceClient.Create(context.TODO(), obj, metav1.CreateOptions{})
	if err != nil {
		errorMsg := fmt.Sprintf("Create resource failure %v", err)
		log.Printf("Create resource failure: %v", err)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	// 3. Reply resonse.
	response := map[string]interface{}{
		"status":  "success",
		"message": fmt.Sprintf("Resource %s created successfully", result.GetName()),
		"details": map[string]interface{}{
			"name":      result.GetName(),
			"namespace": obj.GetNamespace(),
			"kind":      obj.GetKind(),
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
