package controller

import (
	"context"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"

	examplealpha1 "github.com/flyer103/minimalist-operator/pkg/apis/example/v1alpha1"
	crclientset "github.com/flyer103/minimalist-operator/pkg/clients/clientset/versioned"
	crinformer "github.com/flyer103/minimalist-operator/pkg/clients/informers/externalversions/example/v1alpha1"
)

type Controller struct {
	crClient crclientset.Interface
	crSynced cache.InformerSynced
}

func NewController(crClient crclientset.Interface, crInformer crinformer.ExampleInformer) *Controller {
	controller := &Controller{
		crClient: crClient,
		crSynced: crInformer.Informer().HasSynced,
	}

	klog.InfoS("Set up event handlers.")
	crInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    controller.add,
		UpdateFunc: controller.update,
		DeleteFunc: controller.delete,
	})

	return controller
}

func (c *Controller) Run(stopCh <-chan struct{}) error {
	klog.InfoS("Run controller.")

	klog.InfoS("Wait for informer cache to sync.")
	if ok := cache.WaitForCacheSync(stopCh, c.crSynced); !ok {
		return errors.New("Failed to wait for caches to sync.")
	}

	klog.InfoS("Start worker.")
	<-stopCh
	klog.InfoS("Shut down.")

	return nil
}

func (c *Controller) add(obj interface{}) {
	klog.InfoS("Receive ADD Event.")

	exampleObj, ok := obj.(*examplealpha1.Example)
	if !ok {
		klog.Errorf("Failed to type assert object: %v", obj)
		return
	}
	klog.InfoS("obj", "namespace", exampleObj.Namespace, "name", exampleObj.Name)

	ret := exampleObj.DeepCopy()
	ret.Status.Message = "Received In ADD"
	_, err := c.crClient.ProductV1alpha1().Examples(ret.Namespace).UpdateStatus(context.TODO(), ret, metav1.UpdateOptions{})
	if err != nil {
		klog.ErrorS(err, "Failed to update status", "namespace", ret.Namespace, "name", ret.Name)
		return
	}
	klog.InfoS("Update Status.", "namespace", ret.Namespace, "name", ret.Name)
}

func (c *Controller) update(old, new interface{}) {
	klog.InfoS("Receive UPDATE Event.")
}

func (c *Controller) delete(obj interface{}) {
	klog.InfoS("Receive DELETE Event.")
}
