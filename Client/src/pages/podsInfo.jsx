import React, { useEffect, useState } from "react";
import { Card, Button, Form, Input, Table, Modal, message, Space } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import MyUpload from "../components/myUpload";
import axios from "axios";
import "./pagesCSS.css";
import { useFetch } from "../utils/useFetch";

function PodsInfo() {
  const { TextArea } = Input;
  const [yamlContent, setYamlContent] = useState("");
  const [isShow, setIsShow] = useState(false);
  const [myForm] = Form.useForm();
  const {
    data: tableData,
    loading,
    error,
  } = useFetch("http://localhost:8080/k8spodlist.html");
  const [current, setCurrent] = useState(1);
  const [pageSize, setPageSize] = useState(10);

  const [pageOption, setPageOption] = useState({
    pageNo: 1,
    pageSize: 10,
  });

  const dataSource = { tableData };
  const paginationProps = {
    current: pageOption.pageNo,
    pageSize: pageOption.pageSize,
    total: dataSource.length,
    onChange: (current, size) => paginationChange(current, size),
  };

  const paginationChange = async (current, size) => {
    setPageOption({
      pageNo: current,
      pageSize: size,
    });
  };

  //Deploy yaml
  const sendDataToBackend = async (yamlContent) => {
    try {
      console.log("Sending YAML content:", yamlContent);
      const response = await axios.post(
        "http://localhost:8080/api/deploypod",
        yamlContent,
        {
          headers: {
            "Content-Type": "application/yaml",
          },
        }
      );
      return response.data;
    } catch (error) {
      console.error("Error sending yaml:", error);
      if (error.response) {
        console.error("Error response:", error.response.data);
        console.error("Error status:", error.response.status);
      }
      throw error;
    }
  };

  // 统一的提交处理函数
  const handleSubmit = async () => {
    try {
      // 直接使用yamlContent
      const response = await sendDataToBackend(yamlContent);
      console.log("Success:", response);
      message.success("Post Yaml to backend api Successfully!");
      setIsShow(false);
      // 清空YAML内容
      setYamlContent("");
    } catch (error) {
      console.error("Error:", error);
      if (error.response) {
        message.error(
          `Failed to post Yaml to backend api: ${
            error.response.data?.error || "Unknown error"
          }`
        );
      } else {
        message.error("Failed to post Yaml to backend api: Network error");
      }
    }
  };

  const columns = [
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
  ];

  return (
    <div>
      <Card
        style={{ borderColor: "#ac48ebff" }}
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
        <Form form={myForm} style={{ overflow: "auto" }}>
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
            rowKey={(record) => record.key}
          />
        )}
      </Card>
      <Modal
        title="Add Yaml content:"
        onCancel={() => {
          setIsShow(false);
        }}
        maskClosable={false}
        open={isShow}
        onOk={handleSubmit}
        cancelText="CANCEL"
        okText="SEND YAML"
        width={800}
        footer={null}
        // footer={[
        //   <Button key="cancel" onClick={() => setIsShow(false)}>
        //     CANCEL
        //   </Button>,
        //   <Button key="submit" type="primary" onClick={handleSubmit}>
        //     SEND YAML
        //   </Button>,
        // ]}
      >
        <Form form={myForm} labelCol={{ span: 3 }}>
          <Form.Item label="">
            <div style={{ display: "flex", flexDirection: "column" }}>
              <TextArea
                style={{ width: "100%" }}
                rows={25}
                value={yamlContent}
                onChange={(e) => setYamlContent(e.target.value)}
                maxLength={100000}
              />
              <div style={{ marginTop: 16 }} >
                <MyUpload onFileContentChange={setYamlContent} />
                <span></span>
                <div style={{float: "right", display: "inline-block"}}>
                <Button key="cancel" onClick={() => setIsShow(false)} style={{marginRight: 20}}>
                  CANCEL
                </Button>
                <Button key="submit" type="primary" onClick={handleSubmit}>
                  SEND YAML
                </Button>
                </div>
              </div>
            </div>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}

export default PodsInfo;
