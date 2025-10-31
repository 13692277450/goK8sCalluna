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
          <Route path="namespaceInfoModify" element={<NamespaceInfoModify />} />
          <Route path="course_menu" element={<CourseManagment />} />
        </Routes>
      </Mylayout>   
      </div>
  );
}
export default App;
