import React, { useEffect, useState } from "react";
import { Card, Button, Form, Input, Table, Modal, message, Flex, Spin } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import MyUpload from "../components/myUpload";
import axios from "axios";
// import { get } from "../utils/request";
import "../css/pagesCSS.css";
import { useFetch } from "../utils/useFetch";
import { API_BASIC_URL } from "../utils/useFetch";


function NamespaceInfoModify() {
  const { TextArea } = Input;
  const [yamlContent, setYamlContent] = useState(`# pls refer to the example below:
apiVersion: v1
kind: Namespace
metadata:
  name: example-namespace
  labels:
    app: example
    environment: production
  annotations:
    description: "This namespace is for example application"
    managed-by: "zulu-agent"`);
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
    pageNo: 1, //分页序号
    pageSize: 10, //每页显示的条数
  });

  /**
   * 分页配置
   */
  const dataSource = { tableData };
  const paginationProps = {
    current: pageOption.pageNo,
    pageSize: pageOption.pageSize,
    total: dataSource.length,
    onChange: (current, size) => paginationChange(current, size),
  };
  /**
   * 分页改变
   */
  const paginationChange = async (current, size) => {
    setPageOption({
      pageNo: current, //当前所在页面
      pageSize: size, //当前所在页面数据数量
    });
  };
  const sendDataToBackend = async (yamlContent) => {
    try {
      console.log("Sending YAML content:", yamlContent);
      // 使用API_BASIC_URL构建完整的URL
      const url = `${API_BASIC_URL}/deploynamespace`;
      const response = await axios.post(url, yamlContent, {
        headers: {
          "Content-Type": "application/yaml",
        },
      });
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
      }}
  const columns = [
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
  ];
  return (
    <div>
      <Card
        style={{
          borderColor: "#ac48ebff",
          marginTop: 10,
          marginRight: 10,
          overflow: "auto",
        }}
        title="NameSpace Details"
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
        <span><h4>Loading namespace data, pls wait...  </h4>  <Spin /></span>
        </Flex>
      </div>
        ) : error ? (
          <div>Error loading namespace data: {error.message}</div>
        ) : tableData && tableData.length === 0 ? (
          <div style={{textAlign: 'center', padding: '20px'}}>No data available</div>
        ) : (
          <Table
            columns={columns}
            dataSource={tableData}
            pagination={paginationProps}
            rowKey={(record) => record.key}
            rowClassName={() => "custom-row-line-purple"}
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
              <div style={{ marginTop: 16 }}>
                <MyUpload onFileContentChange={setYamlContent} />
                <span></span>
                <div style={{ float: "right", display: "inline-block" }}>
                  <Button
                    key="cancel"
                    onClick={() => setIsShow(false)}
                    style={{ marginRight: 20 }}
                  >
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

export default NamespaceInfoModify;
