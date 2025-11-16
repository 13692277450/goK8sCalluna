import React from "react";
import { Col, Row, Card, Form, Input, Button, message } from "antd";
import reactLog from "../assets/react.svg";
import { useNavigate } from "react-router-dom";

function LogOut() {
    const navigate = useNavigate();
  return (
    <Row>
      <Col
        md={{
          span: 8,
          push: 8,
        }}
        xs={{
          span: 22,
          push: 1,
        }}
        sm={{
          span: 12,
          push: 6,
        }}
      >
        <img src={reactLog} style={{ margin: 20 }}></img>
        <Card
          title="K8S MANAGE SYSTEM CENTER LOGIN"

          style={{ margin: 20, 
            backgroundColor: "lightblue" ,
        }}
        headStyle={{textAlign:'center', fontSize:20, color:"green"}}
        >
          <Form
            labelCol={{
              md: {
                span: 6,
              },
            }}
            onFinish={(n)=>{{
                console.log('Submit', n);
                message.info('Login successful')
                navigate('/admin/system_status')
            }}}
          >
            <Form.Item
              label="Username"
              name="username"
              rules={[
                { required: true, message: "Please input your username(anything)!" },
              ]}
            >
              <Input placeholder="username" />
            </Form.Item>
            <Form.Item
              label="Password"
              name="password"
              rules={[
                { required: true, message: "Please input your password(anything)!" },
              ]}
            >
              <Input.Password placeholder="password" />
            </Form.Item>
            <Form.Item style={{alignContent:'center', alignItems:'center'}}>
              <Button type="primary" htmlType="submit" style={{display:'block', margin:'0 auto'}}>
                Login
              </Button>
            </Form.Item>
          </Form>
          <h1>Logout</h1>
          <p>You have been logged out</p>
          <p>
            Click <a href="/login">here</a> to login again.
          </p>
          <p>
            Click <a href="/">here</a> to go back home.
          </p>
        </Card>
      </Col>
    </Row>
  );
}

export default LogOut;
