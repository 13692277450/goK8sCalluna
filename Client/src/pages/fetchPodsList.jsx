import React from "react";
import { useFetch } from "../utils/useFetch";

function FetchPodsList() {
  const url = "http://localhost:8080/k8spodlist.html";
  const { data, loading, error } = useFetch(url);
  console.log("data:", data);
return (
  <div>
    <h2>Fetch Pods List</h2>
  </div>
);
}
export default FetchPodsList;