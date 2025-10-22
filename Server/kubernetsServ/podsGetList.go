package kubernetsServ

import (
	"context"
	"fmt"
	"log"

	"gok8s/models"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// var Kubeconfig *string
// var Config *rest.Config
// var Clientset *kubernetes.Clientset

// // embeded: kubeconfig: ./kubeconfig
// var kubeconfigEmbed embed.FS

// func K8sInit() {

// 	// if home := homedir.HomeDir(); home != "" {
// 	// 	Kubeconfig = flag.String("Kubeconfig", filepath.Join(home, ".kube", "config"), "./config/Kubeconfig")
// 	// } else {
// 	Kubeconfig = flag.String("kubeconfig", "kubeconfig", "")
// 	// }
// 	flag.Parse()

// 	fmt.Println("Kubeconfig:", *Kubeconfig)

// 	Config, err := clientcmd.BuildConfigFromFlags("", *Kubeconfig)
// 	if err != nil {
// 		fmt.Println("Config errorL: ", err)
// 		panic(err.Error())
// 	}
// 	fmt.Println("config: ", Config)

// 	Config.APIPath = "api"
// 	Config.GroupVersion = &corev1.SchemeGroupVersion
// 	Config.NegotiatedSerializer = scheme.Codecs

// 	restClient, err := rest.RESTClientFor(Config)
// 	if err != nil {
// 		fmt.Println("restClient error: ", err)
// 		panic(err.Error())
// 	}
// 	// Clientset, err := kubernetes.NewForConfig(config)
// 	// configInCluster, err := rest.InClusterConfig()

// 	result := &corev1.PodList{}

// 	fmt.Println("result: ===========", result)
// 	err = restClient.Get().
// 		Namespace("whalebase").
// 		Resource("pods").
// 		VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).
// 		Do(context.TODO()).
// 		Into(result)
// 	fmt.Println("result: ", result)
// 	for _, d := range result.Items {
// 		fmt.Println("range: error:: ", err)

// 		fmt.Printf("namespace:%v \t name:%v \t status:%+v\n", d.Namespace, d.Name, d.Status.Phase)
// 	}
// }

func GetK8sPods() []models.PodInfo {
	// Kubeconfig = flag.String("kubeconfig", "kubeconfig", "")
	// Config, err := clientcmd.BuildConfigFromFlags("", *Kubeconfig)
	// if err != nil {
	// 	fmt.Println("BuildConfigFromFlags error: ", err)
	// 	panic(err.Error())
	// }
	// Clientset, err := kubernetes.NewForConfig(Config)
	// if err != nil {
	// 	fmt.Println("Clientset  error: ", err)
	// 	log.Fatal(err)
	// }
	// 获取 Pod 列表
	pods, err := Clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("List Pods error: ", err)
		log.Fatal(err)
	}

	fmt.Println("Pod List: ")
	var PodInfos []models.PodInfo
	for _, pod := range pods.Items {
		podInfo := models.PodInfo{
			Name:      pod.Name,
			Status:    string(pod.Status.Phase),
			NodeName:  pod.Spec.NodeName,
			HostIP:    pod.Status.HostIP,
			PodIP:     pod.Status.PodIP,
			StartTime: pod.Status.StartTime.String(),
			Namespace: pod.Namespace,
		}
		PodInfos = append(PodInfos, podInfo)
		fmt.Printf("名称：%s，状态：%s, NodeName: %s, HostIP: %s, PodIP: %s, StartTime: %s, NameSpace: %s \n",
			pod.Name, pod.Status.Phase, pod.Spec.NodeName, pod.Status.HostIP, pod.Status.PodIP, pod.Status.StartTime, pod.Namespace)
	}
	models.SetPods(PodInfos)
	return PodInfos
	//ok
	// var ns, label, field string
	// flag.StringVar(&ns, "namespace", "", "namespace")
	// flag.StringVar(&label, "l", "", "Label selector")
	// flag.StringVar(&field, "f", "", "Field selector")

	// api := Clientset.CoreV1()
	// // setup list options
	// listOptions := metav1.ListOptions{
	// 	LabelSelector: label,
	// 	FieldSelector: field,
	// }
	// pvcs, err := api.PersistentVolumeClaims(ns).List(context.TODO(), listOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Pvc List: ", pvcs)
}

//fmt.Println("Pvcs: ", pvcs)

//======================
// deployment := &appsv1.Deployment{
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name: "example-deployment",
// 	},
// 	Spec: appsv1.DeploymentSpec{
// 		Replicas: int32Ptr(3),
// 		Selector: &metav1.LabelSelector{
// 			MatchLabels: map[string]string{"app": "example-app"},
// 		},
// 		Template: corev1.PodTemplateSpec{
// 			ObjectMeta: metav1.ObjectMeta{
// 				Labels: map[string]string{"app": "example-app"},
// 			},
// 			Spec: corev1.PodSpec{
// 				Containers: []corev1.Container{
// 					{
// 						Name:  "example-container",
// 						Image: "nginx:latest",
// 					},
// 				},
// 			},
// 		},
// 	},
// }
// result, err := Clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
// if err != nil {
// 	fmt.Println("Create deployment error: ", err)
// 	log.Fatal(err)
// }
// fmt.Printf("Deployment %q 已创建\n", result.GetObjectMeta().GetName())
//}

//func int32Ptr(i int32) *int32 { return &i }
