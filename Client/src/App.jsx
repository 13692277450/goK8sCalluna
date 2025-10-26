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
import Clock from "./pages/SysncClock";
import SynscClock from "./pages/SysncClock";
import FetchPodsList from "./pages/fetchPodsList";
import PodsLogs from "./pages/podsLogs";
import ClusterInfo from "./pages/clusterInfo";


function App() {
  const [count, setCount] = useState(0);
 // FetchPodsList();
  return (    
    <div>
      <Mylayout>
        <Routes>
          <Route path="system_status" element={<SystemStatus />} />
          <Route path="cluster_info" element={<ClusterInfo />} />
          <Route path="podsInfo" element={<PodsInfo />} />
          <Route path="pvcsInfo" element={<PVCsInfo />} />
          <Route path="resourcesInfo" element={<ResourcesInfo />} />
          <Route path="podsLogs" element={<PodsLogs />} />
          <Route path="SynscClock" element={<SynscClock />} />
          <Route path="course_menu" element={<CourseManagment />} />
        </Routes>
      </Mylayout>   
      </div>
  );
}
export default App;
