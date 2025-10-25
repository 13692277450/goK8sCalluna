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
	// 处理请求
	if r.Method != http.MethodPost {
		http.Error(w, "不支持的请求方法", http.StatusMethodNotAllowed)
		return
	}

	// 1. 读取请求体
	yamlData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("读取请求体失败: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 检查YAML内容是否为空
	if len(yamlData) == 0 {
		http.Error(w, "YAML内容不能为空", http.StatusBadRequest)
		return
	}

	// 2. 处理数据
	// 使用正确的配置变量
	log.Printf("尝试使用配置创建动态客户端: %v", kubernetsServ.ConfigR)
	dynamicClient, err := dynamic.NewForConfig(kubernetsServ.ConfigR)
	if err != nil {
		log.Printf("创建动态客户端失败，错误详情: %v", err)
		http.Error(w, fmt.Sprintf("创建Kubernetes客户端失败: %v", err), http.StatusInternalServerError)
		return
	}

	// 解码 YAML 为 Unstructured 对象
	//yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	decode := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

	obj := &unstructured.Unstructured{}
	_, _, err = decode.Decode(yamlData, nil, obj)
	if err != nil {
		log.Printf("YAML解码失败，错误详情: %v", err)
		http.Error(w, fmt.Sprintf("YAML解析失败: %v", err), http.StatusBadRequest)
		return
	}
	log.Printf("成功d解码YAML，资源类型: %s", obj.GetKind())

	// 验证必需的字段
	if obj.GetKind() == "" {
		http.Error(w, "YAML中缺少必需的kind字段", http.StatusBadRequest)
		return
	}

	// 获取资源类型并转换为复数形式
	kind := obj.GetKind()
	resource := strings.ToLower(kind)
	if !strings.HasSuffix(resource, "s") {
		resource += "s" // 例如 Deployment -> deployments
	}

	// 获取命名空间，默认为default
	namespace := obj.GetNamespace()
	if namespace == "" {
		namespace = "default"
	}

	// 创建资源客户端
	gvr := obj.GroupVersionKind()
	log.Printf("准备创建资源，类型: %s, 命名空间: %s", resource, namespace)
	resourceClient := dynamicClient.Resource(gvr.GroupVersion().WithResource(resource)).Namespace(namespace)

	// 创建资源
	log.Printf("发送创建请求...")
	result, err := resourceClient.Create(context.TODO(), obj, metav1.CreateOptions{})
	if err != nil {
		errorMsg := fmt.Sprintf("创建资源失败: %v", err)
		log.Printf("创建资源失败，错误详情: %v", err)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	// 3. 返回响应
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
