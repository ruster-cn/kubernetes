package main

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	toolCache "k8s.io/client-go/tools/cache"
)

func main() {
	pod := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "test1"}}
	key, err := toolCache.MetaNamespaceKeyFunc(pod)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(key)
}
