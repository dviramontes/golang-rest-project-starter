import React, { useState } from "react"

interface Alert{
  ID: string,
  Text: string,
}

interface AlertsProps {
  alerts: Alert[],
  deleteFn: Function,
  upvoteAlert: Function,
}

interface AletsInputProps {
  fn: Function,
}
const Alerts: React.FC<AlertsProps> =
  ({ alerts = [], deleteFn, upvoteAlert }) => (
    <>
      {alerts.length === 0 ?
        <p>No alerts</p> :
        <ul>
          {
            alerts.map(({ ID, Text}) =>
              (<li key={ID}>
                {`${ID} - ${Text}`}
                <span onClick={() => deleteFn(ID)}>
                  <b><small>{" [x] "}</small></b>
                </span>
                <span onClick={() => upvoteAlert(ID)}>
                  <b><small>{" [^] "}</small></b>
                </span>
              </li>))
          }
        </ul>
      }
    </>)

const AlertInput: React.FC<AletsInputProps> =
  ({ fn }) => {
    const [alert, setAlert] = useState("")
    const enterKeyGuard = ({ keyCode }: React.KeyboardEvent) => {
      if (keyCode === 13) {
        if (isAllCaps(alert)) {
          fn(alert)
        } else {
          window.alert("alert must be all caps")
        }
      }
    }

    return (
      <input type="text"
             onKeyDown={enterKeyGuard}
             onChange={(e) => setAlert(e.target.value)}
             placeholder="enter new alert"/>)

  }

const isAllCaps = (str: string) => str === str.toUpperCase()

export { Alerts, AlertInput, isAllCaps }
