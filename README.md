## **GoK8sCalluna - K8s Manage Tool**

Front and backend full stack design. server code can running on Server or client, because it use .kube\\config file in Golang code with embed file block, but later I should conside kubeconfig was changed and how to get fresh config from server directely.

K8s manage tool build by Golang，React, Vite, Ant-D, Js, CSS, Mysql, K8s, Docker, Ubuntu, Metrics-server, this project plan build an excellent K8s management tool in the next two years, pls kindly support me and join together if you are sinteresting! Thanks. Now, it can get Pods info and show, resources info, log info, deploy pod..., I will continue update.

## **DEMO address:  http://104.168.125.34/ (a little bit slow as server resources limited.)**
The username and password to access Prometheus and Grafana: username: admin, password: 868686msM

<span style="color: #843fa1;">**How to run this application? so easy, 5 steps only:**</span>

1\. Setup your K8S cluster system and metrics server0.8.0, Kafka4.1.0.  
2\. enter into server folder, copy your K8s server config from /roo/.kube/config to Golang code Server\kubeconfig.  
3\. Modify Server\\kubeconfig file, chang server: https://192.168.1.xxx:6443 to your K8s master server if need.  
4\. go run . then can run your backend server.  
5\. Enter into client folder, modify uitils\\useFech.js, modify ip address to your ip address: " API_BASIC_URL = 'http://localhost:8080/api'; //connect to Golang backend api server port. "  
6\. Run npm i then run  "npm run dev".  
7\. Then you can access K8s manage tool by browser.

<span style="color: #3598db;">**or DON'T install K8s server, enter into server folder, run go mod tidy then go run . , the backend server wiil start. then enter into client folder, run npm i then run npm run dev, you can access client UI by browser now.**</span>

*2025.11.16, update: Add service information page.

*2025.11.09, update: Add Prometheus and Grafana monitor dashboard on Kubernets servre and put them to K8s Manage Tool UI.

<span style="color: #e67e23;">*2025.11.04, update: share my VPS K8s cluster for this project, just run server and client is enough to trial run, no necessary install K8s cluster, fixed bugs.*</span>

*2025.11.05 update: Setup client and compile server backend as linux code and put on vps, pls access http://104.168.125.34/ for demo.

*2025.10.31 update: Add delete function in pods information page(pls notice for permant delete pod, this code need delete pod's create style including deploy, namespace, cronjob...., the code can identify which create pod type used but be careful if you have other pod in deployment or namespace..., you can modify code you want which one is your prefer chooice.), build function to add namespace function page. will try add kafka function in code at next stage.*


<img width="2559" height="1524" alt="image" src="https://github.com/user-attachments/assets/a7a1173a-e7dc-4282-91e6-9a1f61d8a593" />



<img width="2552" height="1281" alt="image" src="https://github.com/user-attachments/assets/f891839c-06c8-4761-a565-d9e60ef9d15b" />


<img width="2559" height="1599" alt="image" src="https://github.com/user-attachments/assets/9320dffe-f3a9-470d-a6ae-dd9abf2ef78d" class="jop-noMdConv">

&nbsp;

*2025.10.27 update: Add pods running status page and setup auto refresh per 5 minutes.*

<img width="2531" height="1524" alt="image" src="https://github.com/user-attachments/assets/e8f585f2-ca53-438d-aa19-d1946bc51318" class="jop-noMdConv">

&nbsp;

*2025.10.26 update: Add k8s cluster server cpu, memory... etc usage collect and monitor from metrics-server ver 0.8.0 in backend server side, update metrics-server data demonstrate on webpage in client UI.*

<img width="2559" height="1232" alt="image" src="https://github.com/user-attachments/assets/030c7a3f-de8a-469e-b2f0-0ea2052f5002" class="jop-noMdConv">

*2025.10.25 update: Add function to deploy Yaml to server auto create resources, pod... etc.*

<img width="2551" height="1531" alt="image" src="https://github.com/user-attachments/assets/803239f1-b03b-4675-9595-aba9c6746865" class="jop-noMdConv"> <img width="2554" height="1517" alt="image" src="https://github.com/user-attachments/assets/ac4667dd-1153-4c90-a44d-8130764de3a0" class="jop-noMdConv">

&nbsp;

<img width="2530" height="1459" alt="image" src="https://github.com/user-attachments/assets/22ce423a-363f-463a-a7b3-c8bd4b732ed6" class="jop-noMdConv">

&nbsp;

<img width="2543" height="1527" alt="image" src="https://github.com/user-attachments/assets/d5d23600-72cb-4ea6-a7b9-eedd71f2e63b" class="jop-noMdConv"> <img width="2518" height="1504" alt="image" src="https://github.com/user-attachments/assets/5fa5161a-b23e-42fa-874f-1c2b6ad9e0f6" class="jop-noMdConv">

License: MIT License.
