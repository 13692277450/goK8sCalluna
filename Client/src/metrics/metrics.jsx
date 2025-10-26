import React from "react";
import { Stat, Progress, Placeholder } from "rsuite";
import { useFetch } from "../utils/useFetch";
import { API_BASIC_URL } from "../utils/useFetch";
import { Alert, Card } from "antd";

// 格式化内存显示
const formatMemory = (bytes) => {
  if (!bytes && bytes !== 0) return "0 B";

  const units = ["B", "KB", "MB", "GB", "TB"];
  let size = bytes;
  let unitIndex = 0;

  while (size >= 1024 && unitIndex < units.length - 1) {
    size /= 1024;
    unitIndex++;
  }

  return `${size.toFixed(2)} ${units[unitIndex]}`;
};

// MetricsDashboard - 使用rsuit组件的函数式结构
const MetricsDashboard = () => {
  // 从API获取节点数据 - 使用正确的API URL
  const { data: rawNodeData, loading, error } = useFetch(`${"metrics/nodes"}`);

  // 处理API返回的数据 - 确保数据格式正确
  const processNodeData = (rawData) => {
    if (!rawData) return [];

    // 如果数据是数组，直接返回
    if (Array.isArray(rawData)) {
      return rawData;
    }

    // 如果数据是对象，尝试转换为数组
    if (typeof rawData === "object" && rawData !== null) {
      // 假设API返回的是 { node1: {...}, node2: {...} } 格式
      // 转换为 [{node_name: 'node1', ...}, {node_name: 'node2', ...}]
      return Object.keys(rawData).map((key) => ({
        node_name: key,
        ...rawData[key],
      }));
    }

    return [];
  };

  const nodeData = processNodeData(rawNodeData);

  // 加载状态组件
  const LoadingPlaceholder = () => (
    <Placeholder.Grid columns={1} active>
      <Placeholder.Paragraph style={{ height: 120 }} />
      <Placeholder.Paragraph style={{ height: 120 }} />
      <Placeholder.Paragraph style={{ height: 120 }} />
    </Placeholder.Grid>
  );

  // 错误提示组件
  const ErrorAlert = ({ error }) => (
    <Card
      type="inner"
      title="Error"
      style={{ background: "#f9e5e5", border: "1px solid #f5222d" }}
    >
      <Alert
        message="Error Loading Metrics"
        description={error.message}
        type="error"
        showIcon
      />
    </Card>
  );

  // 根据节点数据创建统计卡片 - 添加可视化组件
  const renderNodeStats = (node) => {
    if (!node) return null;

    // 为不同节点分配不同颜色
    const getNodeColor = (nodeName) => {
      const colors = ["#1890ff", "#52c41a", "#faad14", "#f5222d", "#722ed1"];
      const hash = nodeName
        .split("")
        .reduce((acc, char) => acc + char.charCodeAt(0), 0);
      return colors[hash % colors.length];
    };

    // 确保CPU值在0-1范围内
    const cpuValue = Math.min(1, Math.max(0, node.cpu || 0)) * 100;
    const memoryValue = Math.min(1, Math.max(0, node.memory || 0)) * 100;

      return (
      <Card 
        key={node.node_name}
        bordered={true}
        style={{ 
          textAlign: "center",
          alignSelf: "center",
          alignItems: "center",
          marginBottom: '20px',
          width: '100%',
          minWidth: '300px'
        }}
      >
        <div style={{
          display: 'flex',
          flexDirection: 'column',
          padding: '16px'
        }}>
          <div style={{
            display: 'flex',
            justifyContent: 'space-between',
            marginBottom: '16px',
            marginRight:'80px',
            color:"darkblue",

          }}>
            <Stat.Label>NODE NAME:  {node.node_name}</Stat.Label>
          </div>
          
          <div style={{
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            marginBottom: '8px'
          }}>
            <span>CPU USAGE</span>
            <div style={{ 
              color:'blue',
              fontSize: '16px',

              position: 'relative',
              width: '100%',
              maxWidth: '100px',
              aspectRatio: '1/1'  // 保持正方形
            }}>
              <Progress.Circle
                trailColor="gray"
                percent={cpuValue}
                value={cpuValue}
                width="100%"
                strokeColor={getNodeColor(node.node_name)}
                strokeWidth={8}
                trailWidth={8}
                max={100}
                showInfo={false}
              />
              <div style={{
                position: 'absolute',
                top: '50%',
                left: '50%',
                transform: 'translate(-50%, -50%)',
                fontSize: '14px'
              }}>
                {(cpuValue).toFixed(0)}%
              </div>
            </div>
          </div>
          
          <div style={{
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center'
          }}>
            <span>MEMORY USAGE</span>
            <Stat.Value>{formatMemory(node.memory)}</Stat.Value>
          </div>
        </div>
      </Card>
    );
  };

  return (
    <div style={{ padding: "24px", background: "#f0f2f5" }}>
      <h1 style={{ textAlign: "center", marginBottom: "32px", color:'blueviolet'}}>
        KUBERNETS METRICS DASHBOARD
      </h1>

      {/* 加载状态和错误处理 */}
      {loading && <LoadingPlaceholder />}
      {error && <ErrorAlert error={error} />}

      {/* 显示节点指标 - 使用CSS Grid实现三列排列 */}
      {!loading && !error && nodeData && (
        <div style={{
          display: 'grid',
          gridTemplateColumns: 'repeat(auto-fill, minmax(300px, 1fr))',
          gap: '60px',
          width: '100%'
        }}>
          {nodeData.map(renderNodeStats)}
        </div>
      )}
    </div>
  );
};

export default MetricsDashboard;
