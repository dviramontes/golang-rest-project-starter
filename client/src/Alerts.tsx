import React from "react"

interface Alert{
  ID: string,
  Text: string,
}

interface AlertsProps {
  alerts: Alert[],
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

export default Alerts
