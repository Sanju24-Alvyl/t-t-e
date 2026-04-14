import React, { useState } from "react";

function App() {
  const [message, setMessage] = useState("");

  const callBackend = async () => {
    const res = await fetch("http://35.200.221.163:3000/api/hello");
    const data = await res.json();
    setMessage(data.message);
  };

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h1>Two-Tier App</h1>
      <button onClick={callBackend}>Call Backend</button>
      <p>{message}</p>
    </div>
  );
}

export default App;