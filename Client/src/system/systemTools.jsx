import React from "react";
import { DownOutlined, SettingOutlined } from "@ant-design/icons";
import { Dropdown, Space, Divider, Flex, Input,Cascader } from "antd";
import NodesMenu from "./nodesMen";
// import { Kafka } from "kafkajs";
const { TextArea } = Input;
const optionsNodesMenu = [
  {
    value: '1',
    label: 'K8s-master01',
  },
  {
    value: '2',
    label: 'K8s-node01',

  },
    {
    value: '3',
    label: 'K8s-node02',
  },
];

const onChange = (e) => {
  console.log("Change:", e.target.value);
};
const onClick = ({ key }) => {
  run();
 // message.info(`Click on item ${key}`);
};
const items = [
  {
    key: "1",
    label: "Read Kafka Messages",
    disabled: true,
  },
  {
    type: "divider",
  },
  {
    key: "2",
    label: "Profile",
    extra: "⌘P",
  },
  {
    key: "3",
    label: "Billing",
    extra: "⌘B",
  },
  {
    key: "4",
    label: "Settings",
    icon: <SettingOutlined />,
    extra: "⌘S",
  },
];
const SystemTools = () => (
  <div>
    <div style={{ float: "right", marginRight: 10 ,fontWeight: 'bold'}}>
      <NodesMenu />
      </div>
    <div style={{ float: "right", marginRight: 10 ,fontWeight: 'bold'}}>
      <Dropdown menu={{ items }}>
        <a onClick={(e) => e.preventDefault()}>
          <Space>
            SYSTEM TOOLS
            <DownOutlined />
          </Space>
        </a>
      </Dropdown>
    </div>
    <br />
    <Divider style={{ borderColor: "#38f30fff", width: "100%" }} />
    <br />
    <div style={{ height: "100vh" }}>
      <Flex vertical style={{ height: "100%" }}>
        <TextArea
          showCount
          maxLength={100000}
          onChange={onChange}
          //placeholder="disable resize"
          style={{
            height: "100%",
            width: "100%",
            border: "1px solid #38f30fff",
            overflow: "auto",
          }}
        />
      </Flex>
    </div>
  </div>
);
export default SystemTools;
