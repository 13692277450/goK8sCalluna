/**
 * 解析Kubernetes日志格式为JSON
 * @param {string} logText - 原始日志文本
 * @returns {Array} 解析后的JSON对象数组
 */
export function parseKubeLog(logText) {
  const logs = [];
  const lines = logText.trim().split('\n');
  
  // 正则表达式匹配日志格式
  const logRegex = /^(\w{3} \d{2} \d{2}:\d{2}:\d{2}) (\S+) (\S+): (\w)(\d{4} \d{2}:\d{2}:\d{2}\.\d{6}) +(\d+) (\S+:\d+)] "([^"]+)" err="([^"]+)"/;
  
  lines.forEach(line => {
    const match = line.match(logRegex);
    if (match) {
      logs.push({
        timestamp: match[1],
        node: match[2],
        component: match[3],
        level: match[4],
        detailTime: match[5],
        pid: match[6],
        fileLine: match[7],
        message: match[8],
        error: match[9]
      });
    }
  });
  
  return logs;
}

/**
 * 将解析后的日志转换为JSON字符串
 * @param {string} logText - 原始日志文本
 * @returns {string} JSON字符串
 */
export function LogToJson(logText) {
  const parsedLogs = parseKubeLog(logText);
  return JSON.stringify(parsedLogs, null, 2);
}