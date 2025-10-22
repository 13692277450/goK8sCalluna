import React, { useEffect, useState } from "react";
import { Card, Button, Form, Input, Table, Modal, message } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import MyUpload from "../components/myUpload";
import axios from "axios";
import { get } from "../utils/request";
import "./pagesCSS.css";
import { useFetch } from "../utils/useFetch";


function ResourcesInfo() {
  const [isShow, setIsShow] = useState(false);
  const [myForm] = Form.useForm();
  const { data: tableData, loading, error } = useFetch("http://localhost:8080/resourcesInfo.html");
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
              title: "Kind",
              dataIndex: "Kind",
              key: "Kind",
            },
            {
              title: "SingularName",
              dataIndex: "SingularName",
              key: "SingularName",
              
            },
            {
              title: "DeepCopyGroup",
              dataIndex: "DeepCopyGroup",
              key: "DeepCopyGroup",
            },
            {
              title: "DeepCopyName",
              dataIndex: "DeepCopyName",
              key: "DeepCopyName",
            },
            {
              title: "Verbs",
              dataIndex: "Verbs",
              key: "Verbs",
            },
            {
              title: "Namespaced",
              dataIndex: "Namespaced",
              key: "Namespaced",
            },
            {
              title: "Group",
              dataIndex: "Group",
              key: "Group",
            },
              {
              title: "Version",
              dataIndex: "Version",
              key: "Version",
            },
               {
              title: "ShortNames",
              dataIndex: "ShortNames",
              key: "ShortNames",
            },
             {
              title: "StorageVersionHash",
              dataIndex: "StorageVersionHash",
              key: "StorageVersionHash",
            },
          ]



  return (
      <div>
      <Card
        title="Resources Details"
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
          <Form.Item label="Resources Name:">
            <Input placeholder="Pls input resource name:" />
          </Form.Item>
        </Form>
        {loading ? (
          <div>Loading resources data...</div>
        ) : error ? (
          <div>Error loading resources data: {error.message}</div>
        ) : (
          <Table style={{overflow: 'auto'}}
            columns={columns}
            dataSource={tableData}
            pagination={paginationProps}
            rowKey={record => record.key}
          />
        )}
</Card>
      <Modal
        title="ADD"
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
            message.success("Add sucess");

            setIsShow(false);
          }}
        >
          <Form.Item
            label="Name:"
            name="name"
            rules={[
              {
                required: true,
                message: "Pls input name:",
              },
              {
                max: 20,
                message: "Maximum characters is limite to 20",
                whitespace: true,
              },
            ]}
          >
            <Input placeholder="Pls input name:" />
          </Form.Item>
          <Form.Item label="name:" name="desc">
            <Input placeholder="pls input  singularName:" />
          </Form.Item>
          <Form.Item label="Kind">
            <MyUpload />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}

export default ResourcesInfo;
