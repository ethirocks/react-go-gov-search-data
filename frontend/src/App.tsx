import React from "react";
import DataDisplay from "./components/DataDisplay";

const App: React.FC = () => {
  return (
    <div className="App">
      <h1>Government Data Viewer</h1>
      <DataDisplay />
    </div>
  );
};

export default App;
