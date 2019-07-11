import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'
import { isAllCaps } from './Alerts';

it('renders without crashing', () => {
  const div = document.createElement('div')
  ReactDOM.render(<App />, div)
  ReactDOM.unmountComponentAtNode(div)
})

test('all caps', () => {
  expect(isAllCaps("FOOBAR")).toBe(true)
  expect(isAllCaps("foobar")).toBe(false)
})
