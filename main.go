/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
	// "regexp"
)

var redis_sts_1 = `
{
	"apiVersion": "v1",
	"kind": "ConfigMap",
	"metadata": {
		"name": "redis-cluster"
	},
	"data": {
		"update-node.sh": "#!/bin/sh\nREDIS_NODES=\"/data/nodes.conf\"\nsed -i -e \"/myself/ s/[0-9]\\{1,3\\}\\.[0-9]\\{1,3\\}\\.[0-9]\\{1,3\\}\\.[0-9]\\{1,3\\}/${POD_IP}/\" ${REDIS_NODES}\nexec \"$@\"\n",
		"redis.conf": "cluster-enabled yes\ncluster-require-full-coverage no\ncluster-node-timeout 15000\ncluster-config-file /data/nodes.conf\ncluster-migration-barrier 1\nappendonly yes\nprotected-mode no"
	}
}
`

var redis_sts_2 = `
{
	{
		"apiVersion": "apps/v1",
		"kind": "StatefulSet",
		"metadata": {
			"name": "redis-cluster"
		},
		"spec": {
			"serviceName": "redis-cluster",
			"replicas": 6,
			"selector": {
				"matchLabels": {
					"app": "redis-cluster"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "redis-cluster"
					}
				},
				"spec": {
					"containers": [
						{
							"name": "redis",
							"image": "redis:5.0.1-alpine",
							"ports": [
								{
									"containerPort": 6379,
									"name": "client"
								},
								{
									"containerPort": 16379,
									"name": "gossip"
								}
							],
							"command": [
								"/conf/update-node.sh",
								"redis-server",
								"/conf/redis.conf"
							],
							"env": [
								{
									"name": "POD_IP",
									"valueFrom": {
										"fieldRef": {
											"fieldPath": "status.podIP"
										}
									}
								}
							],
							"volumeMounts": [
								{
									"name": "conf",
									"mountPath": "/conf",
									"readOnly": false
								},
								{
									"name": "data",
									"mountPath": "/data",
									"readOnly": false
								}
							]
						}
					],
					"volumes": [
						{
							"name": "conf",
							"configMap": {
								"name": "redis-cluster",
								"defaultMode": 493
							}
						}
					]
				}
			},
			"volumeClaimTemplates": [
				{
					"metadata": {
						"name": "data"
					},
					"spec": {
						"accessModes": [
							"ReadWriteOnce"
						],
						"resources": {
							"requests": {
								"storage": "1Gi"
							}
						}
					}
				}
			]
		}
	}
`

var redis_svc = `
{
	"apiVersion": "v1",
	"kind": "Service",
	"metadata": {
		"name": "redis-cluster"
	},
	"spec": {
		"type": "LoadBalancer",
		"ports": [
			{
				"port": 6379,
				"targetPort": 6379,
				"name": "client"
			},
			{
				"port": 16379,
				"targetPort": 16379,
				"name": "gossip"
			}
		],
		"selector": {
			"app": "redis-cluster"
		}
	}
}
`

func main() {
	var kubeconfig *string

	var home string

	// Get kube config file: $Home/.kube/config
	if home = homeDir(); home != "" {
		fmt.Printf("%s \n", home)
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	//??
	flag.Parse()

	//TODO read config from yaml files and setup the k8s cluster
	// yaml.NewYAMLOrJSONDecoder()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset via config file
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message

		// pathList := strings.Split(home, "/")
		// namespace := pathList[len(pathList)-1]

		//Use default namespace
		namespace := "default"

		//TODO use helm, the package manager for kubernetes
		//TODO get pod name
		//Ref: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#pod-v1-core

		//hardcode Test first node
		pod := "redis-cluster-0"

		_, err = clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				pod, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		}

		time.Sleep(10 * time.Second)
	}
}

func homeDir() string {
	fmt.Printf("homeDir:\n")
	if h := os.Getenv("HOME"); h != "" {
		fmt.Printf(h)
		fmt.Printf("\n")
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
