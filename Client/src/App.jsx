import { useState } from "react";
import "antd/dist/reset.css";
// for date-picker i18n
import "dayjs/locale/zh-cn";
import Mylayout from "./components/Mylayout";
import { Routes, Route, Router } from "react-router-dom";
import SystemStatus from "./pages/systemStatus";
import ClusterInfo from "./pages/k8sStatus";
import CourseManagment from "./pages/courseMenu";
import ResourcesInfo from "./pages/resourcesInfo";
import PodsInfo from "./pages/podsInfo";
import { Content } from "antd/es/layout/layout";
import Clock from "./pages/SysncClock";
import SynscClock from "./pages/SysncClock";
import FetchPodsList from "./pages/fetchPodsList";
import PodsLogs from "./pages/podsLogs";


function App() {
  const [count, setCount] = useState(0);
 // FetchPodsList();
  return (    
    <div>
      <Mylayout>
        <Routes>
          <Route path="system_status" element={<SystemStatus />} />
          <Route path="clusterinfo" element={<ClusterInfo />} />
          <Route path="podsInfo" element={<PodsInfo />} />
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
