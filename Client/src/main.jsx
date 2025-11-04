import { createRoot } from "react-dom/client";
import "antd/dist/reset.css";

import "./index.css";
import App from "./App.jsx";
import { ConfigProvider } from "antd";
import { HashRouter as Router, Route, Routes } from "react-router-dom";
import zhCN from "antd/locale/zh_CN";
import LogOut from "./pages/logOut.jsx";

createRoot(document.getElementById("root")).render(
  <Router>
    <ConfigProvider locale={zhCN}>
      <Routes>
        <Route path="/" element={<LogOut />} />
        <Route path="/admin/*" element={<App />} />
      </Routes>
    </ConfigProvider>
  </Router>
);
