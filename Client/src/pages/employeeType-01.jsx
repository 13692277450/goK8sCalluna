import React, { useEffect, useState } from "react";
import { Card, Button, Form, Input, Table, Modal, message } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import MyUpload from "../components/myUpload";
import axios from "axios";
import { get } from "../utils/request";
import "../css/pagesCSS.css";

function EmployeeType1() {
  const [isShow, setIsShow] = useState(false);
  const [myForm] = Form.useForm();
  const [tableData, setTableData] = useState();

  //   axios.get('/api/getData').then(_d => console.log(_d))
  get("/getData").then((_d) => setTableData(_d.data));

  return (
      <div>
      <Card
        title="员工成绩"
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
        <Form form={myForm}>
          <Form.Item label="姓名：">
            <Input placeholder="请输入姓名" />
          </Form.Item>
        </Form>
        <Table
          dataSource={tableData}
        
          columns={[
            {
              title: "序号3",
              dataIndex: "ids",
              key: "serial",
              width: 180,

              render: (n, m, k) => {
                return <span>{k + 1}</span>;
              },
            },
            {
              title: "姓名",
              dataIndex: "name",
              key: "name",
            },
            {
              //   img:<img src={"img"}/>,
              title: "照片",
              // dataIndex: "img",
              key: "img",
              width: 300,
              render: (n, m, k) => {
                // return console.log("nimg",tableData)
                return <img src={n.img} className="listImg" />; //console.log("value: ", n.img)
              },
            },
            {
              title: "成绩",
              dataIndex: "score",
              key: "score",
            },
          ]}
        ></Table>
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
            label="姓名:"
            name="name"
            rules={[
              {
                required: true,
                message: "请输入姓名",
              },
              {
                max: 10,
                message: "最多输入10个字符",
                whitespace: true,
              },
            ]}
          >
            <Input placeholder="请输入姓名" />
          </Form.Item>
          <Form.Item label="成绩：" name="desc">
            <Input placeholder="请输入成绩" />
          </Form.Item>
          <Form.Item label="照片">
            <MyUpload />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}

export default EmployeeType1;
