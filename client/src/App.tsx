import React from 'react'
import useAxios from 'axios-hooks'
import { AlertInput, Alerts } from './Alerts'
import './App.css'

const App: React.FC = () => {
  const [{ data: getAlerts, loading: getLoading, error: getError }, refetch] = useAxios({
    url: "http://localhost:4000/api/alerts/",
  })

  const [{  error: postError }, execPost] = useAxios({
    url: "http://localhost:4000/api/alerts/",
    method: "POST",
  }, { manual: true })

  const submitAlert = (text: string) => {
    execPost({
      data: {
        text,
      }
    })
  }

  const [{ error: deleteError }, execDelete] = useAxios({
    url: "http://localhost:4000/api/alerts/",
    method: "DELETE",
  }, { manual: true })

  const deleteAlert = (id: number) => {
    execDelete({
      params: {
        id,
      }
    })
  }

  const someError = () => [getError, postError, deleteError].find(e => e)

  return (
    <div className="App">
      <header className="App-header">
        <h1><u>{getLoading ? "Loading..." : "Alerts"}</u></h1>
      </header>
      <button id="refresh-btn" onClick={refetch}>refresh</button>
      <div className="content">
        {!getLoading && <Alerts alerts={getAlerts} deleteFn={deleteAlert}/>}
        <AlertInput fn={submitAlert}/>
      </div>
      <pre style={ { color: "red"}}>
        {someError() && someError().message }
      </pre>
    </div>
  )
}

export default App
