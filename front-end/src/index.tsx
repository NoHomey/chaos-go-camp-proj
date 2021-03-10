import User from "./service/User"

console.log("RUNNING")
const srvc = new User()
srvc.SignIn({
    email: "ivo@test.com",
    password: "Pass123asd"
})
.OnError(err => console.log("error", err))
.OnFail(err => console.error("fail", err))
.OnResult(usr => {
    console.info("result", usr)
})
.Handle()

setTimeout(() => {
    srvc.Access()
    .OnError(err => console.log("error", err))
    .OnFail(err => console.error("fail", err))
    .OnResult(data => console.log("REFRESH", data))
    .Handle()
}, 1000)

/*import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import reportWebVitals from './reportWebVitals';

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();*/
