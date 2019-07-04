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
      </header>
      <div className="content">
        { alerts.length === 0 ?
            <p>No alerts</p> :
            <ul>
              {
                alerts.map((a, i) =>
                    <li key={i}>{`${i} - ${a}`}</li>)
              }
            </ul>
        }
      </div>
    </div>
  );
}

export default App;
