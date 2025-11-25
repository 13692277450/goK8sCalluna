import React from "react";
import { useFetch } from "../utils/useFetch";
import { Card, Spin, Tooltip } from "antd";

const cardHeaderStyle = {
  background: "#e6e6fa", // 淡紫色背景
  color: "#068f34ff", // 青色文字
  fontWeight: "bold",
  borderBottom: "1px solid #ccc",
};

const cardStyle = {
  border: "2px solid #008000", // 绿色边框
  borderRadius: "8px", // 圆角边框
  marginBottom: "20px",
};

const withTooltip = (text, placement = "right") => (
  <Tooltip title={text} placement={placement}>
    {text}
  </Tooltip>
);

// 安全获取嵌套属性值的函数
const getSafeValue = (obj, path, defaultValue = null) => {
  if (!obj) return defaultValue;
  return path.split(".").reduce((o, p) => {
    if (o === null || o === undefined) return defaultValue;
    return o[p];
  }, obj);
};

function NetworkInfo() {
  const url = `networkinfo`;
  const { data: networkData, loading, error } = useFetch(url);

  if (loading) {
    return (
      <div style={{ padding: "20px", textAlign: "center", color:"orange"}}>
        <h4>Loading Network information data, please wait... </h4>
        <Spin size="large" />
      </div>
    );
  }

  if (error) {
    return (
      <div style={{ padding: "20px" }}>
        <Card title="Error" style={cardStyle}>
          <p>Failed to load network information: {error}</p>
        </Card>
      </div>
    );
  }

  // 渲染单个Pod卡片
  const renderPodCard = (pod, index) => (
    <Card
      key={`pod-${index}`}
      size="small"
      title={withTooltip(
        ` ${`${getSafeValue(pod, "namespace", "N/A")}/${getSafeValue(
          pod,
          "name",
          "N/A"
        )}`}`
      )}
      //title=  {`${getSafeValue(pod, 'namespace', 'N/A')}/${getSafeValue(pod, 'name', 'N/A')}`}
      headStyle={cardHeaderStyle}
      style={{
        ...cardStyle,
        flex: "1 0 22%", // 每行四个卡片，自动平分空间
        minWidth: "200px", // 最小宽度，避免在小屏幕上过于拥挤
        marginRight: "20px",
      }}
    >
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "8px",
          padding: "10px",
          wordWrap: "break-word",
          whiteSpace: "normal",
          overflowWrap: "break-word",
          wordBreak: "break-word",
        }}
      >
        <div>
          <strong>Status:</strong> {getSafeValue(pod, "status", "N/A")}
        </div>
        <div>
          <strong>Pod IP:</strong> {getSafeValue(pod, "podIP", "N/A")}
        </div>
        <div>
          <strong>Host IP:</strong> {getSafeValue(pod, "hostIP", "N/A")}
        </div>
        <div>
          <strong>Node:</strong> {getSafeValue(pod, "nodeName", "N/A")}
        </div>
        <div>
          <strong>Containers:</strong>
          <ul
            style={{ marginTop: "4px", paddingLeft: "20px", marginBottom: 0 }}
          >
            {pod.containers?.map((container, cIndex) => (
              <li key={`container-${cIndex}`} style={{ marginBottom: "4px" }}>
                {getSafeValue(container, "name", "N/A")} (
                {getSafeValue(container, "image", "N/A")})
              </li>
            )) || <li>N/A</li>}
          </ul>
        </div>
      </div>
    </Card>
  );

  // 渲染单个Service卡片
  const renderServiceCard = (service, index) => (
    <Card
      key={`service-${index}`}
      size="small"
      title={`${getSafeValue(service, "namespace", "N/A")}/${getSafeValue(
        service,
        "name",
        "N/A"
      )}`}
      headStyle={cardHeaderStyle}
      style={{
        ...cardStyle,
        flex: "1 0 22%",
        minWidth: "200px",
        marginRight: "20px",
      }}
    >
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "8px",
          padding: "10px",
          wordWrap: "break-word",
          whiteSpace: "normal",
          overflowWrap: "break-word",
          wordBreak: "break-word",
        }}
      >
        <div>
          <strong>Type:</strong> {getSafeValue(service, "type", "N/A")}
        </div>
        <div>
          <strong>Cluster IP:</strong>{" "}
          {getSafeValue(service, "clusterIP", "N/A")}
        </div>
        <div>
          <strong>Ports:</strong>
          <ul
            style={{
              marginTop: "4px",
              paddingLeft: "20px",
              marginBottom: 0,
              wordWrap: "break-word",
              whiteSpace: "normal",
            }}
          >
            {service.ports?.map((port, pIndex) => (
              <li
                key={`port-${pIndex}`}
                style={{
                  marginBottom: "4px",
                  wordWrap: "break-word",
                  whiteSpace: "normal",
                }}
              >
                {getSafeValue(port, "name", "")}{" "}
                {getSafeValue(port, "port", "N/A")}/
                {getSafeValue(port, "protocol", "N/A")}
                {getSafeValue(port, "nodePort") &&
                  ` (NodePort: ${getSafeValue(port, "nodePort", "N/A")})`}
              </li>
            )) || <li>N/A</li>}
          </ul>
        </div>
        <div>
          <strong>Selector:</strong>
          <ul
            style={{
              marginTop: "4px",
              paddingLeft: "20px",
              marginBottom: 0,
              wordWrap: "break-word",
              whiteSpace: "normal",
            }}
          >
            {service.selector ? (
              Object.entries(service.selector).map(([key, value], sIndex) => (
                <li
                  key={`selector-${sIndex}`}
                  style={{
                    marginBottom: "4px",
                    wordWrap: "break-word",
                    whiteSpace: "normal",
                  }}
                >
                  {key}={value}
                </li>
              ))
            ) : (
              <li>N/A</li>
            )}
          </ul>
        </div>
      </div>
    </Card>
  );

  // 渲染单个Node卡片
  const renderNodeCard = (node, index) => (
    <Card
      key={`node-${index}`}
      size="small"
      title={getSafeValue(node, "name", "N/A")}
      headStyle={cardHeaderStyle}
      style={{
        ...cardStyle,
        flex: "1 0 22%",
        minWidth: "200px",
        marginRight: "20px",
      }}
    >
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "8px",
          padding: "10px",
          wordWrap: "break-word",
          whiteSpace: "normal",
          overflowWrap: "break-word",
          wordBreak: "break-word",
        }}
      >
        <div>
          <strong>Internal IP:</strong>{" "}
          {getSafeValue(node, "internalIP", "N/A")}
        </div>
        <div>
          <strong>External IP:</strong>{" "}
          {getSafeValue(node, "externalIP", "N/A")}
        </div>
        <div>
          <strong>Hostname:</strong> {getSafeValue(node, "hostname", "N/A")}
        </div>
        <div>
          <strong>OS:</strong> {getSafeValue(node, "os", "N/A")}
        </div>
        <div>
          <strong>Kubelet Version:</strong>{" "}
          {getSafeValue(node, "kubeletVersion", "N/A")}
        </div>
        <div>
          <strong>Container Runtime:</strong>{" "}
          {getSafeValue(node, "containerRuntime", "N/A")}
        </div>
      </div>
    </Card>
  );

  // 渲染单个CNI网络策略卡片
  const renderNetworkPolicyCard = (policy, index) => (
    <Card
      key={`policy-${index}`}
      size="small"
      title={`${getSafeValue(policy, "name", "N/A")} (${getSafeValue(
        policy,
        "namespace",
        "N/A"
      )})`}
      headStyle={cardHeaderStyle}
      style={{
        ...cardStyle,
        flex: "1 0 22%",
        minWidth: "200px",
        marginRight: "20px",
      }}
    >
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "8px",
          padding: "10px",
          wordWrap: "break-word",
          whiteSpace: "normal",
          overflowWrap: "break-word",
          wordBreak: "break-word",
        }}
      >
        <div>
          <strong>Policy Types:</strong>{" "}
          {getSafeValue(policy, "policyTypes", []).join(", ") || "N/A"}
        </div>
      </div>
    </Card>
  );

  // 渲染CNI网络插件卡片
  const renderNetworkPluginsCard = () => (
    <Card
      size="small"
      title="CNI Information"
      headStyle={cardHeaderStyle}
      style={{
        ...cardStyle,
        flex: "1 0 22%",
        minWidth: "200px",
        marginRight: "20px",
      }}
    >
      <div
        style={{
          padding: "10px",
          wordWrap: "break-word",
          whiteSpace: "normal",
        }}
      >
        <ul
          style={{
            margin: 0,
            paddingLeft: "20px",
            wordWrap: "break-word",
            whiteSpace: "normal",
          }}
        >
          {networkData.cniInfo.cniConfig && (
            <div
              style={{
                wordWrap: "break-word",
                whiteSpace: "normal",
              }}
            >
              <h4>
                <strong>CNI Config:</strong>{" "}
                <li>{networkData.cniInfo.cniConfig} </li>
              </h4>
            </div>
          )}
          {networkData?.cniInfo?.networkPlugins?.map((plugin, index) => (
            <div>
              <h4
                key={`plugin-${index}`}
                style={{
                  wordWrap: "break-word",
                  whiteSpace: "normal",
                }}
              >
                <strong>Network Plugins: </strong>
              </h4>
              <li>{plugin}</li>
            </div>
          )) || <li>N/A</li>}
        </ul>
      </div>
    </Card>
  );

  return (
    <div style={{ padding: "20px" }}>
      <h2 style={{ marginBottom: "20px", color: "#333" }}>
        Kubernetes Network Information
      </h2>

      {/* 所有卡片的主容器 - 自动换行，每行四个 */}
      <div
        style={{
          display: "flex",
          flexWrap: "wrap",
          gap: "20px 0", // 垂直间距20px，水平间距0（通过marginRight控制）
          width: "100%",
        }}
      >
        {/* Pod卡片 */}
        {networkData?.pods?.map((pod, index) => renderPodCard(pod, index))}

        {/* Service卡片 */}
        {networkData?.services?.map((service, index) =>
          renderServiceCard(service, index)
        )}

        {/* Node卡片 */}
        {networkData?.nodes?.map((node, index) => renderNodeCard(node, index))}

        {/* CNI相关卡片 */}
        {networkData?.cniInfo && (
          <>
            {/* 网络插件卡片 */}
            {networkData.cniInfo.networkPlugins && renderNetworkPluginsCard()}

            {/* 网络策略卡片 */}
            {networkData.cniInfo.networkPolicies?.map((policy, index) =>
              renderNetworkPolicyCard(policy, index)
            )}
          </>
        )}
      </div>
    </div>
  );
}

export default NetworkInfo;
