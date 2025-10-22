import React, { useEffect, useState } from "react";
import { Card, Button, Form, Input, Table, Modal, message } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import MyUpload from "../components/myUpload";
import axios from "axios";
import { get } from "../utils/request";
import "./pagesCSS.css";
import { useFetch } from "../utils/useFetch";

function SystemStatus(){
  return(
    <>
    {/* <div>
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
            </div>
    */}
    <div> 
      SystemStatus function design in progress....
    </div>
    </>
  )
}
export default SystemStatus

function SystemStatus1() {
  const [isShow, setIsShow] = useState(false);
  const [myForm] = Form.useForm();
  const [tableData, setTableData] = useState();
 const [current, setCurrent] = useState(1);
  // 默认一页展示10条数据
  const [pageSize, setPageSize] = useState(10);
  //   axios.get('/api/getData').then(_d => console.log(_d))
  // get("/getData").then((_d) => setTableData(_d.data));
  const { data } = useFetch("http://localhost:8080/k8spodlist.html");
  setTableData(data);
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
 * pageination change event handler

  */
const paginationChange = async (current, size) => {
  setPageOption({
    pageNo: current, //current page number
    pageSize: size,  //page size

  })
}
const columns=[
            {
              title: "System 1",
              dataIndex: "Name",
              key: "Name",
              width: 180,
              render: (text, record, index) => `${(pageOption.pageNo - 1) * 10 + (index + 1)}`,
            },
            {
              title: "System 2",
              dataIndex: "Status",
              key: "Status",
            },
            {
              //   img:<img src={"img"}/>,
              title: "System 3",
              // dataIndex: "img",
              key: "img",
              width: 300,
              render: (n, m, k) => {
                // return console.log("nimg",tableData)
                return <img src={n.img} className="listImg" />; //console.log("value: ", n.img)
              },
            },
            {
              title: "System 4",
              dataIndex: "score",
              key: "score",
            },
          ]



  return (
      <div>
      <Card
        title="System Information Search"
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
          <Form.Item label="System: ">
            <Input placeholder="Pls input system name:" />
          </Form.Item>
        </Form>
        <Table
   columns={columns}
   dataSource={tableData}
   pagination={paginationProps}
   rowKey={record => record.key}
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

//export default SystemStatus1;
