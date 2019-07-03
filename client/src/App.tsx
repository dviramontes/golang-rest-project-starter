import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './App.css';

const endpoint: string = "http://localhost:4000/api/alerts/"

const App: React.FC = () => {
  const [alerts, setAlerts] = useState([]);

  const getAlerts = async ()  => {
    try {
      const { status, data: newAlerts } = await axios.get(endpoint)
      if (status === 200) {
        setAlerts(alerts.concat(newAlerts))
      }
    } catch (e) {
      throw e;
    }
  }

  useEffect(() => {
    getAlerts()
  }, [])

  return (
      <div className="App">
      <header className="App-header">
        <h1><u>Alerts</u></h1>
        <>
          { alerts.length === 0 ?
            <p>No alerts</p> :
            <ul style={{listStyle: "none"}}>
              { alerts.map((a, i) => <li>{`${i} - ${a}`}</li>) }
            </ul>
          }
        </>
      </header>
    </div>
  );
}

export default App;
