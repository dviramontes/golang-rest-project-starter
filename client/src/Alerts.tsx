import React, { useState } from "react"

interface Alert{
  ID: string,
  Text: string,
}

interface AlertsProps {
  alerts: Alert[],
}

interface AletsInputProps {
  fn: Function,
}
const Alerts: React.FC<AlertsProps> =
  ({ alerts = [] }) => (
    <>
      {alerts.length === 0 ?
        <p>No alerts</p> :
        <ul>
          {
            alerts.map(({ ID, Text}) =>
              (<li key={ID}>{`${ID} - ${Text}`}</li>))
          }
        </ul>
      }
    </>)

const AlertInput: React.FC<AletsInputProps> =
  ({ fn }) => {
    const [alert, setAlert] = useState("")
    const enterKeyGuard = ({ keyCode }: React.KeyboardEvent) => {
      if (keyCode === 13) {
        fn(alert)
      }
    }

    return (
      <input type="text"
             onKeyDown={enterKeyGuard}
             onChange={(e) => setAlert(e.target.value)}
             placeholder="enter new alert"/>)

  }
export { Alerts, AlertInput }
