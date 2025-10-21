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

   Run on each worker:
   sudo kubeadm join <MASTER_IP>:6443 --token <TOKEN> \
  --discovery-token-ca-cert-hash sha256:<HASH>


   To join master node we need following details:
   - Token
      to create new: kubeadm token create
      to use existing: kubeadm token list
      <img width="940" height="113" alt="image" src="https://github.com/user-attachments/assets/0af63f23-0a57-4aee-b4ae-37cecdd7b32b" />

   - MASTER_IP
       ip a | grep inet **OR** hostname -I
       <img width="940" height="86" alt="image" src="https://github.com/user-attachments/assets/70040745-c9d4-4f5f-bb79-71adfb0b6880" />

   - HASH
       openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | \
       openssl rsa -pubin -outform der 2>/dev/null | \
       openssl dgst -sha256 -hex | \
       sed 's/^.* //'
      <img width="940" height="128" alt="image" src="https://github.com/user-attachments/assets/14ba1ce0-a679-4f7f-8d7e-41f895a9dbd7" />

      Once Joined below are screenshot to show master and worked nodes
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

**Issues Encountered and Resoultion**
1. ✅ Issue Encountered:
   •	Legacy repo apt.kubernetes.io failed with 404 for both kubernetes-jammy and kubernetes-xenial.
   ✅ Resolution:
   Used new official repo from pkgs.k8s.io:
   curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.29/deb/Release.key | \
   gpg --dearmor | sudo tee /etc/apt/keyrings/kubernetes-apt-keyring.gpg > /dev/null
