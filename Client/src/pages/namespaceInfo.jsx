import React, { useEffect, useState } from "react";
import { Card, Button, Form, Input, Table, Modal, message , Flex, Spin} from "antd";
import { PlusOutlined } from "@ant-design/icons";
import MyUpload from "../components/myUpload";
import axios from "axios";
// import { get } from "../utils/request";
import "../css/pagesCSS.css";
import { useFetch } from "../utils/useFetch";



function NamespaceInfo() {
  const [isShow, setIsShow] = useState(false);
  const [myForm] = Form.useForm();
  const { data: tableData, loading, error } = useFetch("namespaceinfo");
  const [current, setCurrent] = useState(1);
  // 默认一页展示10条数据
  const [pageSize, setPageSize] = useState(10);

/**
 * 分页
 * @param pageOption 
 */
const [pageOption, setPageOption] = useState({
  pageNo: 1,  //分页序号
  pageSize: 10, //每页显示的条数
})

/**
 * 分页配置
 */
const dataSource ={tableData}
const paginationProps = {
  current: pageOption.pageNo,
  pageSize: pageOption.pageSize,
  total: dataSource.length, 
  onChange: (current, size) => paginationChange(current,size)
}
/**
 * 分页改变
  */
const paginationChange = async (current, size) => {
  setPageOption({
    pageNo: current, //当前所在页面
    pageSize: size,  //当前所在页面数据数量
  })
}


const columns=[
            {
              title: "Name",
              dataIndex: "name",
              key: "name",
            },
            {
              title: "Creation Time",
              dataIndex: "creation",
              key: "creation",
            },
            {
              title: "Status",
              dataIndex: "status",
              key: "status",
              
            },
                     ]
  return (
      <div>
      <Card
        style={{borderColor: '#ac48ebff',width:500, height:300, marginBottom: 20, marginRight: 20, overflow: 'auto',
}}
        title="NameSpace Details"
        extra={
          <div>
           
          </div>
        }
      >
       
        {loading ? (
            <div
        style={{
          padding: "20px",
          //backgroundColor: '#f5f5f5',
          borderRadius: "4px",
          color: "#ee8282ff",
          fontStyle: "italic",
          fontSize: "16px",
          textAlign: "center",
        }}
      >
        <Flex align="center" gap="middle">
        <span><h4>Loading nameSpace data, pls wait...  </h4>  <Spin /></span>
        </Flex>
      </div>
        ) : error ? (
          <div>Error loading namespace data: {error.message}</div>
        ) : (
         
          <Table
            columns={columns}
            dataSource={tableData}
            pagination={paginationProps}
            rowKey={record => record.key}
            rowClassName={()=> 'custom-row-line-purple'}

          />
        )}
</Card>
      
    </div>
  );
}

export default NamespaceInfo;
