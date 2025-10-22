import React from "react";
import { useFetch } from "../utils/useFetch";

function fechResources() {
    const url = "http://localhost:8080/resourcesInfo.html"
    const {data, loading, error} = useFetch(url)
    return (
        <div>

        </div>
    )

}
export default fechResources;