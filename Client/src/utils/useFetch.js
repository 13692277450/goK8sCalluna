import { useState, useEffect } from "react";
import axios from "axios";

// 公共API基础URL - 导出为模块变量
export const API_BASIC_URL = 'http://localhost:8080/api'; //connect to Golang backend api server port.

export const useFetch = (endpoint) => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // 拼接完整URL
    const url = `${API_BASIC_URL}/${endpoint}`;
    
    const fetchData = async () => {
      try {
        const res = await axios.get(url);
        setData(res.data);
      } catch (err) {
        console.error(`Error fetching ${url}: ${err}`);
        setError(err);
      } finally {
        setLoading(false);
      }
    };
    
    fetchData();
  }, [endpoint]);

  return { data, loading, error };
};