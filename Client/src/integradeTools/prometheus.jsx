import React from "react";
import "../css/pagesCSS.css";

const Prometheus=()=> {
  const handleNavigate = () => {
    window.open("http://104.168.125.34:30090/classic/status", "_blank");
  };

  return (
    <div style={{color:'green', padding: '20px'}}>
      <h2>Prometheus Service Center</h2>
      <button onClick={handleNavigate} style={{padding: '10px 20px', fontSize: '16px'}}>
        Click To Open Prometheus Service Center
      </button>
    </div>
  );
}

export default Prometheus;
