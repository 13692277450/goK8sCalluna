package kubernetsServ

import (
	"embed"
	"flag"
	"fmt"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	Kubeconfig *string
	Config     *rest.Config // 主配置
	ConfigR    *rest.Config // 资源获取专用配置（保持兼容）
	Clientset  *kubernetes.Clientset
)

// embedded: kubeconfig: ./kubeconfig

var kubeconfigEmbed embed.FS

func K8sConnectionInit() {
	// Initialize client
	// if home := homedir.HomeDir(); home != "" {
	// 	Kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	Kubeconfig = flag.String("kubeconfig", "kubeconfig", "")
	//}
	flag.Parse() // Ensure parsing command line arguments

	var err error
	// Load main configuration
	Config, err = clientcmd.BuildConfigFromFlags("", *Kubeconfig)
	if err != nil {
		fmt.Printf("Failed to load main configuration: %v\n", err)
		panic(err.Error())
	}

	// Initialize resource acquisition configuration (same as main configuration)
	ConfigR = Config

	// Create Clientset
	Clientset, err = kubernetes.NewForConfig(Config)
	if err != nil {
		fmt.Printf("Failed to create Clientset: %v\n", err)
		log.Fatal(err)
	}
}

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
