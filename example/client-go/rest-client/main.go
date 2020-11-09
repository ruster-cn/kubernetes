package main

import (
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	stop := make(chan struct{}, 1)
	config, err := clientcmd.BuildConfigFromFlags("", "./kubeconfig")
	if err != nil {
		panic(err.Error())
	}
	client := kubernetes.NewForConfigOrDie(config)
	factory := informers.NewSharedInformerFactory(client, 1*time.Second)
	nodeInformer := factory.Core().V1().Pods().Informer()
	nodeInformer.AddEventHandler(&cache.ResourceEventHandlerFuncs{})
	go factory.Start(stop)
	cache.WaitForCacheSync(stop, nodeInformer.HasSynced)
	time.Sleep(100 * time.Second)
	close(stop)
}
