import React, { useEffect, useState } from "react";
import { Card, Button, Form, Input, Table, Modal, message, Flex, Spin } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import MyUpload from "../components/myUpload";
import "../css/pagesCSS.css";
import { useFetch } from "../utils/useFetch";

function ServicesInfo() {
  const [isShow, setIsShow] = useState(false);
  const [myForm] = Form.useForm();
  const { data: tableData, loading, error } = useFetch("servicesinfo");
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
    total: dataSource?.length,
    onChange: (current, size) => paginationChange(current, size)
  };

  const paginationChange = async (current, size) => {
    setPageOption({
      pageNo: current,
      pageSize: size,
    });
  };

  // 渲染ports数组的辅助函数
  const renderPortField = (ports, field) => {
    if (!Array.isArray(ports)) return 'Invalid data';
    return ports.map((port, index) => (
      <div style={{width:'80px'}} key={index}>{port[field] || '-'}</div>
    ));
  };

  const columns = [
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
    },
    {
      title: "Namespace",
      dataIndex: "namespace",
      key: "namespace",
    },
    {
      title: "Type",
      dataIndex: "type",
      key: "type",
    },
    {
      title: "ClusterIP",
      dataIndex: "clusterIP",
      key: "clusterIP",
    },
    {
      title: "ExternalIPs",
      dataIndex: "externalIPs",
      key: "externalIPs",
      render: (ips) => ips?.join(', ') || 'None',
    },
  
    {
      title: "Port Name",
      dataIndex: "ports",
      key: "ports.name",
      render: (ports) => renderPortField(ports, 'name'),
      width: '15%',
      align: 'center',
    },
    {
      title: "Target Port",
      dataIndex: "ports",
      key: "ports.targetPort",
      render: (ports) => renderPortField(ports, 'targetPort'),
      width: '15%',
      align: 'center',
    },
    {
      title: "Port Protocol",
      dataIndex: "ports",
      key: "ports.protocol",
      render: (ports) => renderPortField(ports, 'protocol'),
      width: '15%',
      align: 'center',
    },
    {
      title: "Port NodePort",
      dataIndex: "ports",
      key: "ports.nodePort",
      render: (ports) => renderPortField(ports, 'nodePort'),
      width: '15%',
      align: 'center',
    },
      {
      title: "Selector",
      dataIndex: "selector",
      key: "selector",
      render: (selector) => selector ? JSON.stringify(selector) : 'None',
    },
    {
      title: "CreateTime",
      dataIndex: "createTime",
      key: "createTime",
    },
  ];

  return (
    <div>
      <Card
        style={{ borderColor: '#ac48ebff' }}
        title="Services Details"
        extra={
          <div>
            <Button
              type="primary"
              icon={<PlusOutlined />}
              onClick={() => {
                //setIsShow(true);
              }}
            >
              Add
            </Button>
          </div>
        }
      >
        <Form form={myForm} style={{ overflow: 'auto' }}>
          <Form.Item label="Service Name:">
            <Input placeholder="Pls input service name:" />
          </Form.Item>
        </Form>
        {loading ? (
          <div
            style={{
              padding: "20px",
              borderRadius: "4px",
              color: "#98099aff",
              fontStyle: "italic",
              fontSize: "16px",
              textAlign: "center",
            }}
          >
            <Flex align="center" gap="middle">
              <span>
                <h4>Loading services information data, pls wait...</h4>
                <Spin />
              </span>
            </Flex>
          </div>
        ) : error ? (
          <div>Error loading service information: {error.message}</div>
        ) : (
          <Table
            style={{ marginTop: '20px', overflow: 'auto' }}
            columns={columns}
            dataSource={tableData}
            pagination={paginationProps}
            rowKey={(record) => record.name + "-" + record.namespace}
            rowClassName={() => 'custom-row-line-purple'}
           // scroll={{ x: 'max-content' }}
          />
        )}
      </Card>
      <Modal
        title="Add service"
        onCancel={() => {
          setIsShow(false);
        }}
        maskClosable={false}
        open={isShow}
        onOk={() => {
          myForm.submit();
        }}
      >
        <Form
          form={myForm}
          labelCol={{ span: 3 }}
          onFinish={() => {
            message.success("Add service successfully!");
            setIsShow(false);
          }}
        >
          <Form.Item
            label="Name: "
            name="name"
            rules={[
              {
                required: true,
                message: "Pls input service name: ",
              },
              {
                max: 100,
                message: "Maximum characters limited to 100",
                whitespace: true,
              },
            ]}
          >
            <Input placeholder="Pls input service name: " />
          </Form.Item>
          <Form.Item label="Service nName" name="desc">
            <Input placeholder="Pls input service name: " />
          </Form.Item>
          <Form.Item label="Service">
            <MyUpload />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}

export default ServicesInfo;