****Sprint 1****

**Cluster Set Up**

**Master Node**
1. Install EC2 machine using Ubuntu 22.04 (base image ami-03c1f788292172a4e)
<img width="940" height="267" alt="image" src="https://github.com/user-attachments/assets/af3b3452-e847-4ec1-a76d-e35c546e8d0c" />

2. 🧱Run below commands
   sudo apt update && sudo apt upgrade -y
   sudo apt install -y curl bash-completion git apt-transport-https ca-certificates gnupg lsb-release
   sudo swapoff -a
   sudo sed -i '/swap/d' /etc/fstab
   
3. 📦Install containerd (Applicable for all nodes)
   sudo apt install -y containerd
   sudo mkdir -p /etc/containerd
   containerd config default | sudo tee /etc/containerd/config.toml
   sudo sed -i 's/SystemdCgroup = false/SystemdCgroup = true/' /etc/containerd/config.toml
   sudo systemctl restart containerd
   sudo systemctl enable containerd
   
4. 🧠Kernel Modules + Sysctl (Applicable for all nodes)
   cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
   overlay
   br_netfilter
   EOF

   sudo modprobe overlay
   sudo modprobe br_netfilter

   cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
   net.bridge.bridge-nf-call-iptables  = 1
   net.ipv4.ip_forward                 = 1
   net.bridge.bridge-nf-call-ip6tables = 1
   EOF

   sudo sysctl --system


6. 🔧 Install Kubernetes Components (Applicable for all nodes)
   curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.29/deb/Release.key | \
   gpg --dearmor | sudo tee /etc/apt/keyrings/kubernetes-apt-keyring.gpg > /dev/null

   echo "deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.29/deb/ /" | \
   sudo tee /etc/apt/sources.list.d/kubernetes.list

   sudo apt update
   sudo apt install -y kubelet kubeadm kubectl
   sudo apt-mark hold kubelet kubeadm kubectl

 7. 🚦 Master Node Initialization
    sudo kubeadm init --pod-network-cidr=10.244.0.0/16

    Post-init setup:
    mkdir -p $HOME/.kube
    sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
    sudo chown $(id -u):$(id -g) $HOME/.kube/config

8. 🌐 Pod Network (Flannel CNI)
    kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

   AMI created for master is ami-07cd7e2b91f9e1923
   <img width="1625" height="339" alt="image" src="https://github.com/user-attachments/assets/8457d70f-a67c-45c6-aefb-926e94002d09" />


10. Create and Add Worker Node
   Repeat all Steps from 1-6 and the run step 9.

AMI **ami-05540aeb6ec68ed3c** created for node which can we used for master or worker, according will need to run init or join comand after ec2 is launched.
<img width="1606" height="258" alt="image" src="https://github.com/user-attachments/assets/908e3765-7248-4e62-a919-ce742925a9ff" />


   Run on each worker:
   sudo kubeadm join <MASTER_IP>:6443 --token <TOKEN> \
  --discovery-token-ca-cert-hash sha256:<HASH>


   To join master node we need following details:
   - Token
      to create new (run on master): kubeadm token create
     
      to use existing (run on master): kubeadm token list
      <img width="1919" height="105" alt="image" src="https://github.com/user-attachments/assets/4979f04e-18d1-4676-8b4e-ec63c7445ed2" />
      another sample
      <img width="940" height="113" alt="image" src="https://github.com/user-attachments/assets/0af63f23-0a57-4aee-b4ae-37cecdd7b32b" />

   - MASTER_IP
       ip a | grep inet **OR** hostname -I
       <img width="1914" height="103" alt="image" src="https://github.com/user-attachments/assets/db982b2c-2cab-42b6-9b08-8991c15dbd3d" />
       another sample 
       <img width="940" height="86" alt="image" src="https://github.com/user-attachments/assets/70040745-c9d4-4f5f-bb79-71adfb0b6880" />

   - HASH
       openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | \
       openssl rsa -pubin -outform der 2>/dev/null | \
       openssl dgst -sha256 -hex | \
       sed 's/^.* //'
       <img width="1918" height="169" alt="image" src="https://github.com/user-attachments/assets/108ebf5b-365b-4eaa-92d3-0caa04b50e42" />
        another one
      <img width="940" height="128" alt="image" src="https://github.com/user-attachments/assets/14ba1ce0-a679-4f7f-8d7e-41f895a9dbd7" />

      Full Command:
     kubeadm join 172.31.32.27:6443 \
     --token kq12ua.jss372be8i7q1p9k \
     --discovery-token-ca-cert-hash sha256:a834f53ad857245d5a1919ff6049542ff7c4a5e651f003992ddd728f264721ef

      Once Joined below are screenshot to show master and worked nodes
      <img width="1913" height="279" alt="image" src="https://github.com/user-attachments/assets/608bc2f6-abc7-4234-999c-389c5adb5aea" />
      <img width="1919" height="241" alt="image" src="https://github.com/user-attachments/assets/df9f9435-498d-4a63-b649-eb4fbd686e83" />

      <img width="940" height="110" alt="image" src="https://github.com/user-attachments/assets/0c3b592c-7f90-463e-8b4e-b7830b081208" />
      <img width="940" height="249" alt="image" src="https://github.com/user-attachments/assets/6e20192e-fd6a-4114-8b6a-e97b7ee79eeb" />


**Validations/Testing**
List all namespaces
<img width="662" height="162" alt="image" src="https://github.com/user-attachments/assets/584e1abe-88d4-4ab4-8209-90839fcade20" />

list pods in kube-system
<img width="1054" height="203" alt="image" src="https://github.com/user-attachments/assets/7440608c-87ee-46e6-a898-7503299ebeb5" />

kubectl get nodes -o wide
kubectl get pods -n kube-system
kubectl get pods -n kube-flannel
<img width="1915" height="387" alt="image" src="https://github.com/user-attachments/assets/07403e67-46b3-4f73-97dd-0a43b1d23558" />

Prometheus:
Run 
kubectl port-forward svc/monitoring-kube-prometheus-prometheus 9090:9090 -n default

access it on localhost:9090
<img width="1918" height="709" alt="image" src="https://github.com/user-attachments/assets/d17f787c-0fa4-49b8-9044-2aaba48f19bf" />


For testing Grafana and Prometheus
Ran below on local where my key is present
 ssh -i ./Amol-EC2.pem -L 3000:127.0.0.1:3000 ubuntu@52.39.12.74

 It still failed to access from localhost, as 
 Grafana is trying to install plugins like grafana-lokiexplore-app, pyroscope, and exploretraces — but failing due to no internet.

 Fix:
 nano kube-prometheus-values.yaml

 copy below in file
 
 grafana:
  plugins: []

helm upgrade monitoring prometheus-community/kube-prometheus-stack \
  -n default -f kube-prometheus-values.yaml

  or run just below 
  
  helm upgrade monitoring prometheus-community/kube-prometheus-stack \
  -n default \
  --set grafana.plugins={}

**Issues Encountered and Resoultion**
1. ✅ Issue Encountered:
   •	Legacy repo apt.kubernetes.io failed with 404 for both kubernetes-jammy and kubernetes-xenial.
   ✅ Resolution:
   Used new official repo from pkgs.k8s.io:
   curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.29/deb/Release.key | \
   gpg --dearmor | sudo tee /etc/apt/keyrings/kubernetes-apt-keyring.gpg > /dev/null


===
**Install Helm Chart**
   sudo snap install helm --classic
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

   <img width="1918" height="218" alt="image" src="https://github.com/user-attachments/assets/d67617ee-d44e-4d46-a182-f14de98e840b" />
  
   helm repo update
   **install kube-prometheus-stack**
   helm install monitoring prometheus-community/kube-prometheus-stack
   
   **Port-forward Grafana:**
   kubectl port-forward svc/monitoring-grafana 3000:80

   <img width="1919" height="636" alt="image" src="https://github.com/user-attachments/assets/8851095a-c1e7-4e6a-b0ef-7d8983ea5471" />
