import React from "react";
import { useFetch } from "../utils/useFetch";

function FetchPodsList() {
  const url = "k8spodinfo";
  const { data, loading, error } = useFetch(url);
  //console.log("data:", data);
return (
  <div>
  </div>
);
}
export default FetchPodsList;