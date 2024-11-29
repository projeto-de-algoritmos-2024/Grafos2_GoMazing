import React, { useState } from 'react';
import './App.css';

function App() {
  const [maze, setMaze] = useState<string | null>(null);

  const generateMaze = async () => {
    const backendUrl = process.env.REACT_APP_BACKEND_URL || 'http://localhost:8080';
    const response = await fetch(`${backendUrl}/generate-maze`);
    const data = await response.text();
    setMaze(data);
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>GoMazing</h1>
        <button onClick={generateMaze}>Generate Maze</button>
        {maze && <pre>{maze}</pre>}
      </header>
    </div>
  );
}

export default App;