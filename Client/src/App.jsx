import { useState } from "react";
import "antd/dist/reset.css";
// for date-picker i18n
import "dayjs/locale/zh-cn";
import Mylayout from "./components/Mylayout";
import { Routes, Route, Router } from "react-router-dom";
import EmployeeType from "./pages/systemStatus";
import EmployeeList from "./pages/k8sStatus";
import CourseManagment from "./pages/courseMenu";
import PodsListList from "./pages/podsList";
import PodsInfo from "./pages/podsInfo";
import { Content } from "antd/es/layout/layout";
import Clock from "./pages/SysncClock";
import SynscClock from "./pages/SysncClock";
import FetchPodsList from "./pages/fetchPodsList";


function App() {
  const [count, setCount] = useState(0);
  FetchPodsList();
  return (    
    <div>
      <Mylayout>
        <Routes>
          <Route path="system_status" element={<EmployeeType />} />
          <Route path="k8s_status" element={<EmployeeList />} />
          <Route path="podsInfo" element={<PodsInfo />} />
          <Route path="podslist_list" element={<PodsListList />} />
          <Route path="SynscClock" element={<SynscClock />} />
          <Route path="course_menu" element={<CourseManagment />} />
        </Routes>
      </Mylayout>   
      </div>
  );
}
export default App;
