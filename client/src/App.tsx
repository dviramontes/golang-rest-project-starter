import React, { useState } from 'react'
import useAxios from 'axios-hooks'
import { AlertInput, Alerts } from './Alerts'
import './App.css'

const App: React.FC = () => {
  const [alert, setAlert] = useState("")
  const [{data: getAlerts, loading, error}, refetch] = useAxios(
    "http://localhost:4000/api/alerts/all"
  )

  if (error) return <div className="App">error</div>

  return (
    <div className="App">
      <header className="App-header">
        <h1><u>{ loading ? "Loading..." : "Alerts"}</u></h1>
      </header>
      <button id="refresh-btn" onClick={refetch}>refresh</button>
      <div className="content">
        { !loading && <Alerts alerts={getAlerts}/>}
        <AlertInput fn={setAlert}/>
      </div>
    </div>
  )
}

export default App
