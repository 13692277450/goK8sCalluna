import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

export const FetchData = () => {
  return axios.get(`${API_BASE_URL}/deploypod`);
};

export const PostData = (data) => {
  return axios.post(`${API_BASE_URL}/deploypod`, data);
};

export default FetchData; PostData;