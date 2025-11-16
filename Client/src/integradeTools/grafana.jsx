import React from "react";
import "../css/pagesCSS.css";

const Grafana=()=> {
  const handleNavigateKubernets = () => {
    window.open("http://104.168.125.34:30001/d/1235/k8s-dashboard?orgId=1", "_blank");
  };
   const handleNavigateLinux = () => {
    window.open("http://104.168.125.34:30001/d/linux-stats/1-linux-stats-with-node-exporter?orgId=1&refresh=1m", "_blank");
  };
 const handleNavigateDocker = () => {
    window.open("http://104.168.125.34:30001/d/BiigkBu7k/docker-and-system-monitoring-nodeexporter-v1-3-1?orgId=1&refresh=10s", "_blank");
  };
  return (
    <div style={{color:'blue', padding: '20px'}}>
      <h2>Kubenets Cluster Running Dashboard</h2>
      <button onClick={handleNavigateKubernets} style={{padding: '10px 20px', fontSize: '16px', width: '500px'}}>
        Click To Open Kubernets Cluster Running Dashboard
      </button>
      <div>
      <h2> </h2>
      </div>
      <h2>Linux Server Dashboard</h2>
      <button onClick={handleNavigateLinux} style={{padding: '10px 20px', fontSize: '16px', width:'500px'}}>
        Click To Open Linux Server Dashboard
      </button>
       <div>
      <h2> </h2>
      </div>
      <h2>Linux Docker Dashboard</h2>
      <button onClick={handleNavigateDocker} style={{padding: '10px 20px', fontSize: '16px', width:'500px'}}>
        Click To Open Linux Docker Dashboard
      </button>
    </div>
  );
}

export default Grafana;
