import React, { useEffect, useState } from "react";
import { Card, Button, Form, Input, Table, Modal, message } from "antd";
// import { get } from "../utils/request";
import "./pagesCSS.css";

import MetricsPodsDashboard from "../metrics/metricsPods";

function PodsRunningStatus(){

  return(
    <>
    <div> 
           
      <MetricsPodsDashboard />
    </div>
    </>
  )
}
export default PodsRunningStatus
