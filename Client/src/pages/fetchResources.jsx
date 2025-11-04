import React from "react";
import { useFetch } from "../utils/useFetch";

function fechResources() {
    const url = "resourcesinfo"
    const {data, loading, error} = useFetch(url)
    return (
        <div>

        </div>
    )

}
export default fechResources;