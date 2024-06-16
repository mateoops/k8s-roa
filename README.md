# k8s-roa
## Kubernetes Resource Optimization Advisor

<p align="center">
<img loading="lazy" width="400px" src="assets/k8s-roa.jpg" alt="image_name png" />
</p>

---

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=mateoops_k8s-roa&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=mateoops_k8s-roa) 

## About
The Kubernetes Resource Optimization Advisor (k8s-roa) is a tool designed to help optimize the resource allocation for Kubernetes workloads. By analyzing the actual usage patterns of CPU, memory, and other resources, the tool provides recommendations for adjusting resource requests and limits, thereby improving cluster efficiency and reducing costs.

## Example Use Case
A typical use case could involve a web application running in a Kubernetes cluster. The Optimization Advisor would continuously monitor the CPU and memory usage of the application's pods. If it detects that certain pods consistently use less CPU than requested, it would recommend lowering the CPU request, freeing up cluster resources for other workloads. Conversely, if it detects that a pod frequently hits its memory limit, it would recommend increasing the memory limit to prevent performance degradation or crashes. Additionally k8s-roa can constantly monitor the resource utilization on your nodes and recommend amount of nodes based of current load and historical trends.
By implementing the Kubernetes Resource Optimization Advisor, you can achieve a more efficient and cost-effective deployment, ensuring that your Kubernetes workloads run smoothly and are well-provisioned.

## MoSCoW

### Must Have
* collecting usage resources data from Nodes and Pods
* saving data to Prometheus
* reports & recommendations for adjusting resource requests and limits based on usage patterns

### Should Have
* unit tests
* visualization usage and recommendations

### Could Have
* helm chart
* high resources utilization alerts

### Will Not Have
* machine learning model to predict future resource needs
* integration with HPA (based on custom metrics)