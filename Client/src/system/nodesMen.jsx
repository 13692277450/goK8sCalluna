import React from "react";
import { Cascader } from "antd";
import { useFetch } from "../utils/useFetch";
const options = [
  {
    value: "1",
    label: "K8s-master01",
  },
  {
    value: "2",
    label: "K8s-node01",
  },
  {
    value: "3",
    label: "K8s-node02",
  },
];
const onChange = (value) => {
  switch (value) {
    case (value = 1):
      useFetch("systemLogs/kubectlSystemLogs");
      break;
    case (value = 2):
      useFetch("systemLogs/");
      break;
    case (value = 3):
      useFetch("nodes/3");
      break;
  }
  console.log(value);
};
const NodesMenu = () => (
  <Cascader options={options} onChange={onChange} placeholder="Please select" />
);
export default NodesMenu;
