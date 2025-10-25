
2025.10.25 update: Add function to deploy Yaml to server auto create resources, pod... etc.

<img width="2551" height="1531" alt="image" src="https://github.com/user-attachments/assets/803239f1-b03b-4675-9595-aba9c6746865" />

=======================================================================================================================================
Front and backend full stack design. server code can running on Server or client, because it use .kube\config file in Golang code with embed file block, but later I should conside kubeconfig was changed and how to get fresh config from server.

K8s manage tool build by Golang，React, Vite, Ant-D, Js, CSS, Mysql, K8s, Ubuntu, this project plan build an excellent K8s management tool in the next two year, pls kindly support me and join together! Thanks.
Now, it can get Pods info and show, resources info, log info, deploy pod..., I will continue update.

How to run this application? so easy, 4 steps only: 

1. Setup your K8S cluster system.
2. copy .kube\kubeconfig to Server\kubernetsServ\kubeconfig.
3. Modify Server\kubernetsServ\kubeconfig file, chang server: https://192.168.1.211:6443 to your K8s master server if IP address was wrong.
4. go run . then can run your backend server.
5. Enter into client folder, run "npm run dev".
6. Then you can access K8s manage tool by browser.

<img width="2554" height="1517" alt="image" src="https://github.com/user-attachments/assets/ac4667dd-1153-4c90-a44d-8130764de3a0" />

<img width="2530" height="1459" alt="image" src="https://github.com/user-attachments/assets/22ce423a-363f-463a-a7b3-c8bd4b732ed6" />


<img width="2543" height="1527" alt="image" src="https://github.com/user-attachments/assets/d5d23600-72cb-4ea6-a7b9-eedd71f2e63b" />

<img width="2518" height="1504" alt="image" src="https://github.com/user-attachments/assets/5fa5161a-b23e-42fa-874f-1c2b6ad9e0f6" />

