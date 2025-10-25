import React, { useEffect, useState } from "react";
import { Card, Button, Form, Input, Table, Modal, message } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import MyUpload from "../components/myUpload";
import axios from "axios";
// import { get } from "../utils/request";
import "./pagesCSS.css";
import { useFetch } from "../utils/useFetch";
import NodesInfo from "./nodesInfo";
import NamespaceInfo from "./namespaceInfo";

function SystemStatus(){
  return(
    <>
    <div> 
      <NamespaceInfo />
      <NodesInfo />
    </div>
    </>
  )
}
export default SystemStatus
