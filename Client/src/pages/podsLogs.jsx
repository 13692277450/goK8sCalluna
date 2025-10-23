import React from "react";
import { Card, Col, Row } from "antd";
import { useFetch } from "../utils/useFetch";
function PodsLogs() {
  const {
    data: TableData,
    loading,
    error,
  } = useFetch("http://localhost:8080/api/pods/logs");
  const dataSource = TableData;
  
  if (error) {
    console.error('Error fetching logs:', error);
    return <div>Error loading logs</div>;
  }
  if (loading) return <div>Loading data...</div>;
  else {
    console.log('Received data:', dataSource);
    return (
      <div>
        {renderLogCards(dataSource?.data || dataSource)}
        {/* <Row gutter={16}>
          <Col span={8}>
            <Card title="Card title" variant="borderless">
              Card content
            </Card>
          </Col>
          <Col span={8}>
            <Card title="Card title" variant="borderless">
              Card content
            </Card>
          </Col>
          <Col span={8}>
            <Card title="Card title" variant="borderless">
              Card content
            </Card>
          </Col>
        </Row> */}
      </div>
    );
  }
}
export default PodsLogs;


const renderLogCards = (logs) => {
  if (!logs) return null;
  
  try {
    const logText = typeof logs === 'string' ? logs : JSON.stringify(logs);
    const logBlocks = logText.split('TITLE:').filter(block => block.trim());
    const cards = [];
    
    for (const block of logBlocks) {
      const [titlePart, ...contentParts] = block.split('CONTENT:');
      const title1 = titlePart.trim();
      const content = contentParts.join('CONTENT:').trim();
      const lines = logText.split('\n'); // 分割为行数组
      const title = "Pod Name:  "+ lines[0];
      
      if (title && content) {
        cards.push({title, content});
      }
    }
  


    // 每行显示两个卡片
    const rows = [];
    for (let i = 0; i < cards.length; i += 2) {
      rows.push(
        <Row key={i} gutter={16}>
          <Col span={12}>
            <Card 
              title={cards[i].title}
              style={{ 
                marginBottom: 16,
                height: '500px',
                overflow: 'hidden'
              }}
            >
              <div style={{
                height: '450px',
                overflow: 'auto',
                padding: '8px',
                boxSizing: 'border-box'
              }}>
                <pre style={{
                  whiteSpace: 'pre-wrap',
                  margin: 0,
                  wordBreak: 'break-word'
                }}>
                  {cards[i].content}
                </pre>
              </div>
            </Card>
          </Col>
          {cards[i+1] && (
            <Col span={12}>
              <Card 
                title={cards[i+1].title}
                style={{ 
                  marginBottom: 16,
                  height: '500px',
                  overflow: 'hidden'
                }}
              >
              <div style={{
                height: '450px',
                overflow: 'auto',
                padding: '8px',
                boxSizing: 'border-box'
              }}>
                <pre style={{
                  whiteSpace: 'pre-wrap',
                  margin: 0,
                  wordBreak: 'break-word'
                }}>
                  {cards[i+1].content}
                </pre>
              </div>
              </Card>
            </Col>
          )}
        </Row>
      );
    }
    return rows;
  } catch (error) {
    console.error('Error rendering log cards:', error);
    return null;
  }
};

// 使用示例
const LogsDisplay = ({ logsData }) => {
  return (
    <div>
      {renderLogCards(logsData)}
    </div>
  );
};