import React from "react";
import { Stat, Progress, Placeholder } from "rsuite";
import { useFetch } from "../utils/useFetch";
import { useFetchMock } from "../utils/useFetch";
import { API_BASIC_URL } from "../utils/useFetch";
import { Alert, Card, Flex, Spin } from "antd";
import axios from "axios";
//[{"pod_name":"example-deployment55-5549c9ffc-gmcq5","namespace":"default","container":"example-container","cpu":0,"memory":4952064},
// format memory display
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

const MetricsPodsDashboard = () => {
  // 从API获取节点数据 - 使用正确的API URL
  //const { data: rawPodData, loading, error } = useFetchMock("https://mock.presstime.cn/mock/68fee0b65e1f18b9172f3143/example/api/metrics/pods#!method=get");
  const { data: rawPodData, loading, error } = useFetch(`${"metrics/pods"}`);

  // 处理API返回的数据 - 确保数据格式正确
  const processPodData = (rawData) => {
    if (!rawData) return [];

    // 如果数据是数组，直接返回
    if (Array.isArray(rawData)) {
      return rawData;
    }

    // 如果数据是对象，尝试转换为数组
    if (typeof rawData === "object" && rawData !== null) {
      // 假设API返回的是 { Pod1: {...}, Pod2: {...} } 格式
      // 转换为 [{pod_name: 'Pod1', ...}, {pod_name: 'Pod2', ...}]
      return Object.keys(rawData).map((key) => ({
        pod_name: key,
        ...rawData[key],
      }));
    }

    return [];
  };

  // 使用状态管理pod数据
  const [podData, setPodData] = React.useState(processPodData(rawPodData));

  // 每5分钟自动刷新
  React.useEffect(() => {
    let isMounted = true;

    const fetchData = async () => {
      try {
        // const res = await axios.get("https://mock.presstime.cn/mock/68fee0b65e1f18b9172f3143/example/api/metrics/pods#!method=get");
        const res = await axios.get("http://localhost:8080/api/metrics/pods");
        if (isMounted) {
          setPodData(processPodData(res.data));
        }
      } catch (err) {
        console.error(`Error fetching data: ${err}`);
      }
    };

    // 立即执行一次
    fetchData();

    // 设置5分钟(300000毫秒)的定时器
    const interval = setInterval(fetchData, 300000);

    // 在组件卸载时清理定时器
    return () => {
      isMounted = false;
      clearInterval(interval);
    };
  }, []);

  // 当useFetchMock返回的数据变化时更新podData
  React.useEffect(() => {
    setPodData(processPodData(rawPodData));
  }, [rawPodData]);

  // 加载状态组件
  const LoadingPlaceholder = () => (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        padding: "40px",
      }}
    >
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
          <span>
            <h4>Loading pods metrics data, pls wait... </h4> <Spin />
          </span>
        </Flex>
      </div>
    </div>
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
  const renderPodStats = (Pod) => {
    if (!Pod) return null;
    // 确保CPU值在0-1范围内
    const cpuValue = Math.min(1, Math.max(0, Pod.cpu || 0)) * 100;
    const memoryValue = Math.min(1, Math.max(0, Pod.memory || 0)) * 100;
    // 为不同节点分配不同颜色
    const getPodColor = (cpuValue) => {
      const colors = ["#19f432ff", "#eef115ff", "#f70e21ff"];
      // const hash = PodName
      //   .split("")
      //   .reduce((acc, char) => acc + char.charCodeAt(0), 0);
      // return colors[hash % colors.length];
      if (cpuValue > 0.7) return colors[2]; // Red for high CPU usage
      if (cpuValue > 0.4) return colors[1]; // Yellow for medium CPU usage
      return colors[0]; // Green for low CPU usage
    };

    return (
      <Card
        size=""
        key={Pod.pod_name}
        bordered={true}
        style={{
          textAlign: "center",
          alignSelf: "center",
          alignItems: "center",
          marginBottom: "5px",
          width: "100%",
          minWidth: "300px",
          maxWidth: "500px",
          minHeight: "320px",
          maxHeight: "320px",
        }}
      >
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            padding: "16px",
          }}
        >
          <div
            style={{
              display: "flex",
              justifyContent: "space-between",
              marginBottom: "16px",
              marginRight: "80px",
              color: "darkblue",
            }}
          >
            <Stat.Label>POD NAME: {Pod.pod_name}</Stat.Label>
          </div>

          <div
            style={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
              marginBottom: "8px",
              fontWeight: "bold",
            }}
          >
            <span>CPU USAGE</span>
            <div className=".cardcontainer"
              style={{
                color: "blue",
                fontSize: "16px",
                position: "relative",
                width: "100%",
                maxWidth: "100px",
                aspectRatio: "1/1", // 保持正方形
              }}
            >
              <Progress.Circle
                trailColor="lightblue"
                percent={cpuValue}
                value={cpuValue}
                //width="100%"
                strokeColor={getPodColor(Pod.cpu)}
                strokeWidth={8}
                trailWidth={8}
                max={100}
                showInfo={false}
              />
              <div
                style={{
                  position: "absolute",
                  top: "50%",
                  left: "50%",
                  transform: "translate(-50%, -50%)",
                  fontSize: "14px",
                }}
              >
                {cpuValue.toFixed(0)}%
              </div>
            </div>
          </div>
          <div
            style={{
              display: "flex",
              flexDirection: "column",
              width: "100%",
              marginTop: "8px",
              fontFamily: "monospace",
              marginBottom: "5px",
              position: "absolute",
              bottom: "5px",
              alignSelf: "center",
            }}
          >
            <div>
              <span style={{ fontWeight: "bold", color: "blue" }}>
                MEMORY USAGE:{" "}
              </span>
              {formatMemory(Pod.memory)}
            </div>
            <div style={{ marginTop: "4px" }}>
              <span style={{ fontWeight: "bold", color: "blue" }}>
                POD SIZE:{" "}
              </span>
              {formatMemory(Pod.size)}
            </div>
          </div>
        </div>
      </Card>
    );
  };

  return (
    <div style={{ padding: "24px", background: "#f0f2f5" }}>
      <h1
        style={{
          textAlign: "center",
          marginBottom: "32px",
          color: "blueviolet",
        }}
      >
        KUBERNETS PODS DASHBOARD
      </h1>

      {/* 加载状态和错误处理 */}
      {loading && !error && <LoadingPlaceholder />}
      {error && <ErrorAlert error={error} />}

      {/* 显示节点指标 - 使用CSS Grid实现三列排列 */}
      {!loading && !error && podData && (
        <div
          style={{
            display: "grid",
            gridTemplateColumns: "repeat(auto-fill, minmax(300px, 1fr))",
            gap: "20px",
            width: "100%",
          }}
        >
          {podData.map(renderPodStats)}
        </div>
      )}
    </div>
  );
};

export default MetricsPodsDashboard;
