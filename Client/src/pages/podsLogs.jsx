import React from "react";
import { Card, Col, Row, Flex, Spin } from "antd";
import { useFetch } from "../utils/useFetch";

function PodsLogs() {
  const {
    data: TableData,
    loading,
    error,
  } = useFetch('pods/logs')
  
  if (error) {
    console.error('Error fetching logs:', error);
    return <div>Error loading logs</div>;
  }
  if (loading) return   <div
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
        <span><h4>Loading Pods Logs data, pls wait...  </h4>  <Spin /></span>
        </Flex>
      </div>;
  if (!TableData) return <div>No data available</div>;
  console.log('Received data:', TableData);
  return (
    <div>
      {renderLogCards(TableData?.data || TableData)}
    </div>
  );
}

export default PodsLogs;

const renderLogCards = (logs) => {
  if (!logs) return null;
  
  try {
    // 处理JSON格式的日志数据
    let logEntries;
    if (typeof logs === 'string') {
      try {
        logEntries = JSON.parse(logs);
      } catch (e) {
        // 如果解析失败，使用原始字符串
        logEntries = [{podName: "Unknown", logs: logs}];
      }
    } else {
      logEntries = logs;
    }

    // 确保logEntries是数组
    if (!Array.isArray(logEntries)) {
      logEntries = [logEntries];
    }

    const cards = [];
    logEntries.forEach(entry => {
      if (entry.podName && entry.logs) {
        cards.push({
          title: `Pod: ${entry.podName}`,
          content: entry.logs
        });
      }
    });

    // 每行显示两个卡片
    const rows = [];
    for (let i = 0; i < cards.length; i += 2) {
      rows.push(
        <Row key={i} gutter={16}>
          <Col span={12}>
            <Card 
              title={cards[i].title}
              headStyle={{ backgroundColor: '#f0f9f9ff', color: '#869df7ff' }}
              style={{ 
                marginBottom: 16,
                height: '500px',
                overflow: 'hidden',
                borderColor: '#4fe68bff'
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
                headStyle={{ backgroundColor: '#f0f9f9ff', color: '#869df7ff' }}
                style={{ 
                  marginBottom: 16,
                  height: '500px',
                  overflow: 'hidden',
                  borderColor: '#4fe68bff'
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
















// import React from "react";
// import { Card, Col, Row } from "antd";
// import { useFetch } from "../utils/useFetch";
// function PodsLogs() {
//   const {
//     data: TableData,
//     loading,
//     error,
//   } = useFetch('pods/logs')
//   // useFetch("http://localhost:8080/api/pods/logs");
//   const dataSource = TableData;
  
//   if (error) {
//     console.error('Error fetching logs:', error);
//     return <div>Error loading logs</div>;
//   }
//   if (loading) return <div>Loading data...</div>;
//   else {
//     console.log('Received data:', dataSource);
//     return (
//       <div>
//         {renderLogCards(dataSource?.data || dataSource)}
//         {
      
//         }
//       </div>
//     );
//   }
// }
// export default PodsLogs;


// const renderLogCards = (logs) => {
//   if (!logs) return null;
  
//   try {
//     const logText = typeof logs === 'string' ? logs : JSON.stringify(logs);
//     const logBlocks = logText.split('TITLE:').filter(block => block.trim());
//     const cards = [];
    
//     for (const block of logBlocks) {
//       const [titlePart, ...contentParts] = block.split('CONTENT:');
//       const title1 = titlePart.trim();
//       const content = contentParts.join('CONTENT:').trim();
//       const lines = logText.split('\n'); // 分割为行数组
//       const title = "Pod Name:  "+ lines[0].replace("TITLE:","");
      
//       if (title && content) {
//         cards.push({title, content});
//       }
//     }
  


//     // 每行显示两个卡片
//     const rows = [];
//     for (let i = 0; i < cards.length; i += 2) {
//       rows.push(
//         <Row key={i} gutter={16}>
//           <Col span={12}>
//             <Card 
//               title={cards[i].title}
//               headStyle= { { backgroundColor: '#f0f9f9ff', color: '#869df7ff' }}
//               style={{ 
//                 marginBottom: 16,
//                 height: '500px',
//                 overflow: 'hidden',
//                 borderColor: '#4fe68bff'
                
//               }}
//             >
//               <div style={{
//                 height: '450px',
//                 overflow: 'auto',
//                 padding: '8px',
//                 boxSizing: 'border-box'

//               }}>
//                 <pre style={{
//                   whiteSpace: 'pre-wrap',
//                   margin: 0,
//                   wordBreak: 'break-word'
//                 }}>
//                   {cards[i].content}
//                 </pre>
//               </div>
//             </Card>
//           </Col>
//           {cards[i+1] && (
//             <Col span={12}>
//               <Card 
//                 title={cards[i+1].title}
//                 headStyle= { { backgroundColor: '#f0f9f9ff', color: '#869df7ff'  }}
//                 style={{ 
//                   marginBottom: 16,
//                   height: '500px',
//                   overflow: 'hidden',
//                   borderColor: '#4fe68bff',
//                 }}
//               >
//               <div style={{
//                 height: '450px',
//                 overflow: 'auto',
//                 padding: '8px',
//                 boxSizing: 'border-box'
//               }}>
//                 <pre style={{
//                   whiteSpace: 'pre-wrap',
//                   margin: 0,
//                   wordBreak: 'break-word'
//                 }}>
//                   {cards[i+1].content}
//                 </pre>
//               </div>
//               </Card>
//             </Col>
//           )}
//         </Row>
//       );
//     }
//     return rows;
//   } catch (error) {
//     console.error('Error rendering log cards:', error);
//     return null;
//   }
// };

// const LogsDisplay = ({ logsData }) => {
//   return (
//     <div>
//       {renderLogCards(logsData)}
//     </div>
//   );
// };