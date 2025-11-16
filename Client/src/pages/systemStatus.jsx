import React, { useEffect, useState } from "react";
import { Card, Button, Form, Input, Table, Modal, message } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import MyUpload from "../components/myUpload";
import axios from "axios";
// import { get } from "../utils/request";
import "../css/pagesCSS.css";
import MetricsNodesDashboard from "../metrics/metricsNodes";

function SystemStatus(){

  return(
    <>
    <div> 
      <MetricsNodesDashboard />
    </div>
    </>
  )
}
export default SystemStatus
