import React from "react"

interface AlertsProps {
  alerts: string[]
}

const Alerts: React.FC<AlertsProps> =
  ({ alerts }) => (
    <>
      {alerts.length === 0 ?
        <p>No alerts</p> :
        <ul>
          {
            alerts.map((a, i) => <li key={i}>{`${i} - ${a}`}</li>)
          }
        </ul>
      }
    </>)

export default Alerts
