package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/dynamic"
	//"k8s.io/client-go/kubernetes"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/sample-controller/pkg/signals"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	//"github.com/mattkelly/arrive/pkg/controller"

	"github.com/mattkelly/arrive/pkg/apis/arrive/v1alpha1"
	"github.com/mattkelly/arrive/pkg/buildinfo"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	glog.Info("Starting Arrive Controller...")
	glog.Infof("Version: %s", buildinfo.String())
	glog.Infof("Go Version: %s", runtime.Version())

	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	//stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	dyclient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building dynamic client: %s", err.Error())
	}

	resource := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}

	pod, err := dyclient.Resource(resource).Namespace("kube-system").Get("kube-proxy-dgmjl", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v:\n", pod)

	// TEST ONE
	val := "kube-proxy"
	f := v1alpha1.Filter{
		OperandLeft: v1alpha1.Operand{
			ValueFrom: &v1alpha1.ValueReference{
				FieldRef: &v1alpha1.FieldSelector{
					FieldPath: "{.metadata.labels['k8s-app']}",
				},
			},
		},
		Operator: selection.Equals,
		OperandRight: v1alpha1.Operand{
			Value: &val,
		},
	}

	match, err := f.Match(pod)
	if err != nil {
		panic(err)
	}

	fmt.Printf("match: %v\n", match)
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
