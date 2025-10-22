import React, { useEffect, useState } from "react";
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  UploadOutlined,
  UserOutlined,
  ReadOutlined,
  PicLeftOutlined,
  DownOutlined,
  BulbFilled,
  MailOutlined,
  ProfileOutlined,
} from "@ant-design/icons";
import { Button, Layout, Menu, theme, Dropdown, Space, message, Breadcrumb } from "antd";
import logo from "../assets/logo.jpg";
import reactLog from "../assets/react.svg";
import { useNavigate, useLocation } from "react-router-dom";
const { Header, Sider, Content } = Layout;
const items = [
  { key: "userCenter", label: <a>K8S SYSTEM CENTER</a> },
  {
    key: "logOut",
    label: <a>EXIT</a>,
  },
];
const itesMenuData = [
            {
              key: "system_menu",
              label: "SYSTEM MANAGE",
              children: [
                {
                  key: "/admin/system_menu/employee_type",
                  label: "SYSTEM STATUS",
                },
                {
                  key: "/admin/system_menu/employee_list",
                  label: "K8S STATUS",
                },
              ],
            },
            {
              key: "/admin/podslist_menu",
              label: "PODS MANAGE",
              children: [
                {
                  key: "/admin/podslist_menu/podslit_type",
                  label: "PODS LIST",

                },
                {
                  key: "/admin/podslist_menu/resourcesInfo",
                  label: "RESOURCES DETAILS",
                },
              ],
            },
            {
              key: "/admin/course_menu",
              label: "RESOURCE MANAGE",


            },
];

const createNavFn =(key)=>{
  let arrObj =[]
  const demoFn =(arr) =>{
    arr.forEach(n =>{
      const {children, ...info} = n
      arrObj.push(info)
      if(children){
        demoFn(children)
      }
    })
  }
  demoFn(itesMenuData)
  const temp = arrObj.filter(m=>key.includes(m.key))
  if (temp.length > 0){
    return [{label:'HOME',key:'/admin/system_menu/system_status'},...temp]
  } else{} return [];
}

const searchUrlKey = (key) => {
  let arrObj = [];
  const demoFn = (_arr) => {
    _arr.forEach((n) => {
      if (key.includes(n.key)) {
        arrObj.push(n.key);
        if (n.children) {
          demoFn(n.children);
        }
      }
    });
  };
  demoFn(itesMenuData);
  return arrObj;
};
const Mylayout = ({ children }) => {
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();
  const navigate = useNavigate();
  const onClick = ({ key }) => {
    if (key === "logOut") {
      localStorage.removeItem("token");

      navigate("/");
    } else {
      message.info("Function developing in progress...");
    }
  };
  const { pathname } = useLocation();
  let demoItemsArr = searchUrlKey(pathname);
const [navurl,setNavurl] = useState([]);

  useEffect(()=>{
    setNavurl(createNavFn(pathname))
  },[pathname]);
  return (
    <Layout className="ant-layout" style={{ minHeight: "100vh" }}>
      <Sider trigger={null} collapsible collapsed={collapsed}>
        <div className="logoimg">
          <img src={reactLog} alt="" />
        </div>
        <Menu
          theme="dark"
          mode="inline"
          defaultOpenKeys={demoItemsArr}
          defaultSelectedKeys={demoItemsArr}
          onClick={({ key }) => {
            navigate(key);
          }}
          items={[
            {
              key: "system_menu",
              icon: <MailOutlined style={{ color: "green" }} />,
              label: "SYSTEM MANAGE",
              children: [
                {
                  icon: <PicLeftOutlined style={{ color: "orange" }} />,
                  key: "/admin/system_status",
                  label: "SYSTEM STATUS",

                },
                {
                  icon: <PicLeftOutlined style={{ color: "red" }} />,
                  key: "/admin/k8s_status",
                  label: "K8S STATUS",
                },
              ],
            },
            {
              key: "/admin/podslist_menu",
              icon: <ReadOutlined style={{color: "lightgreen"}}/>,
              label: "PODS MANAGE",
              children: [
                {
                  key: "/admin/podsInfo",
                  icon: <ReadOutlined style={{color: "lightgreen"}}/>,
                  label: "PODS INFO",
                },
                {
                  key: "/admin/resourcesInfo",
                  icon: <ReadOutlined style={{color: "lightgreen"}}/>,
                  label: "RESOURCES INFO",
                },
              ],
            },
            {
              key: "/admin/logs_menu",
              icon: <ProfileOutlined style={{color: "cyan"}}/>,
              label: "LOGS MANAGE",
              children: [
                {
                  key: "/admin/podsLogs",
                  icon: <ProfileOutlined style={{color: "cyan"}}/>,
                  label: "LOGS INFO",
                },
                {
                  key: "/admin/resourcesLogs",
                  icon: <ProfileOutlined style={{color: "cyan"}}/>,
                  label: "RESOURCES LOGS",
                },
              ],
            },
              {
              key: "/admin/SynscClock",
              icon: <UserOutlined />,
              label: "CONTAINER MANAGE",

            },
            {
              key: "/admin/course_menu",
              icon: <UploadOutlined />,
              label: "RESOURCE MANAGE",
            },
          ]}
        />
      </Sider>
      <Layout>
        <Header style={{ padding: 0, background: colorBgContainer }}>
          <Button
            type="text"
            icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
            onClick={() => setCollapsed(!collapsed)}
            style={{
              fontSize: "16px",
              width: 64,
              height: 64,
            }}
          />
          <span className="titleDiv">KUBERNETS MANAGE SYSTEM - Ver 0.1</span>
          <Dropdown menu={{ items, onClick }}>
            <a
              onClick={(e) => e.preventDefault()}
              style={{ float: "right", marginRight: 30 }}
            >
              <img
                src={reactLog}
                style={{ width: 20, borderRadius: 15, marginRight: 10 }}
              ></img>
              <Space>
                SYSTEM SETTINGS
                <DownOutlined />
              </Space>
            </a>
          </Dropdown>
        </Header>
        <Content
          style={{
            margin: "24px 16px",
            padding: 24,
            minHeight: 280,
            background: colorBgContainer,
            borderRadius: borderRadiusLG,
          }}
        >
          <Breadcrumb>
          {
            navurl.map(n=>{
              return <Breadcrumb.Item key={n.key}>{n.label}</Breadcrumb.Item>
            })
          }
          </Breadcrumb>
          {children}
        </Content>
      </Layout>
    </Layout>
  );
};
export default Mylayout;
