import { useState } from "react";
import "antd/dist/reset.css";
// for date-picker i18n
import "dayjs/locale/zh-cn";
import Mylayout from "./components/Mylayout";
import { Routes, Route, Router } from "react-router-dom";
import SystemStatus from "./pages/systemStatus";
import CourseManagment from "./pages/courseMenu";
import ResourcesInfo from "./pages/resourcesInfo";
import PodsInfo from "./pages/podsInfo";
import PVCsInfo from "./pages/pvcsInfo";
import { Content } from "antd/es/layout/layout";
import SynscClock from "./pages/SysncClock";
import FetchPodsList from "./pages/fetchPodsList";
import PodsLogs from "./pages/podsLogs";
import PodsRunningStatus from "./pages/podsRunningStatus";
import ClusterStatus from "./pages/clusterInfo";
import NamespaceInfoModify from "./pages/namespaceInfoModify";
import SystemTools from "./system/systemTools";
import Prometheus from "./integradeTools/prometheus";
import Grafana from "./integradeTools/grafana";
import ServicesInfo from "./services/servicesInfo";
import NetworkInfo from "./system/networkInfo";



function App() {
  const [count, setCount] = useState(0);
 // FetchPodsList();
  return (    
    <div>
      <Mylayout>
        <Routes>P
          <Route path="system_status" element={<SystemStatus />} />
          <Route path="cluster_status" element={<ClusterStatus />} />
          <Route path="pods_status" element={<PodsRunningStatus />} />
          <Route path="podsInfo" element={<PodsInfo />} />
          <Route path="pvcsInfo" element={<PVCsInfo />} />
          <Route path="resourcesInfo" element={<ResourcesInfo />} />
          <Route path="podsLogs" element={<PodsLogs />} />
          <Route path="system_tools" element={<SystemTools />} />
          <Route path="namespaceInfoModify" element={<NamespaceInfoModify />} />
          <Route path="system_prometheus" element={<Prometheus />} />
          <Route path="system_grafana" element={<Grafana />} />
          <Route path="services_info" element={<ServicesInfo />} />
          <Route path="network_status" element={<NetworkInfo />} />


        </Routes>
      </Mylayout>   
      </div>
  );
}
export default App;
