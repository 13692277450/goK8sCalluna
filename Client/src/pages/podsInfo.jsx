import React, { useEffect, useState } from "react";
import { Card, Button, Form, Input, Table, Modal, message } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import MyUpload from "../components/myUpload";
import axios from "axios";
import { get } from "../utils/request";
import "./pagesCSS.css";
import { useFetch } from "../utils/useFetch";



function PodsInfo() {
  const [isShow, setIsShow] = useState(false);
  const [myForm] = Form.useForm();
  const { data: tableData, loading, error } = useFetch("http://localhost:8080/k8spodlist.html");
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
              dataIndex: "Name",
              key: "Name",
            },
            {
              title: "Node Name",
              dataIndex: "NodeName",
              key: "NodeName",
            },
            {
              title: "Name Space",
              dataIndex: "Namespace",
              key: "Namespace",
              
            },
            {
              title: "Status",
              dataIndex: "Status",
              key: "Status",
            },
            {
              title: "Host IP",
              dataIndex: "HostIP",
              key: "HostIP",
            },
            {
              title: "Pod IP",
              dataIndex: "PodIP",
              key: "PodIP",
            },
            {
              title: "Start Time",
              dataIndex: "StartTime",
              key: "StartTime",
            },
          ]



  return (
      <div>
      <Card
        title="Pods Details"
        extra={
          <div>
            <Button
              type="primary"
              icon={<PlusOutlined />}
              onClick={() => {
                setIsShow(true);
              }}
            >
              Add
            </Button>
          </div>
        }
      >
        <Form form={myForm} style={{overflow: 'auto'}}>
          <Form.Item label="Pod Name:">
            <Input placeholder="Pls input pod name:" />
          </Form.Item>
        </Form>
        {loading ? (
          <div>Loading pod data...</div>
        ) : error ? (
          <div>Error loading pod data: {error.message}</div>
        ) : (
         
          <Table
            columns={columns}
            dataSource={tableData}
            pagination={paginationProps}
            rowKey={record => record.key}
          />
        )}
</Card>
      <Modal
        title="新增员工成绩"
        onCancel={() => {
          setIsShow(false);
        }}
        maskClosable={false}
        open={isShow}
        onOk={() => {
          myForm.submit();
        }}
      >
        {" "}
        <Form
          form={myForm}
          labelCol={{ span: 3 }}
          onFinish={() => {
            message.success("新增成功");

            setIsShow(false);
          }}
        >
          <Form.Item
            label="Name: "
            name="name"
            rules={[
              {
                required: true,
                message: "Pls input name: ",
              },
              {
                max: 10,
                message: "Maximum characters limited to 10",
                whitespace: true,
              },
            ]}
          >
            <Input placeholder="Pls input name: " />
          </Form.Item>
          <Form.Item label="SingularName" name="desc">
            <Input placeholder="Pls input singularName: " />
          </Form.Item>
          <Form.Item label="Kind">
            <MyUpload />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}

export default PodsInfo;
