import React from "react"

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
  ({ alerts }) => (
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
  ({ fn }) => (
    <input type="text"
           onKeyDown={(key) => console.log(key.keyCode)}
           onChange={(e) => fn(e.target.value)}
           placeholder="enter new alert"/>)

export { Alerts, AlertInput }
