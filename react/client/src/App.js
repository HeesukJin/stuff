import './App.css';
import {useRef} from 'react'

function App() {
  const username = useRef();
  const password = useRef();

  const register = () => {
    console.log('hi')

    fetch('http://localhost:5000/register', {method: 'POST', 
    body: new FormData({username: username.current.value, password: password.current.value})})
      .then((response) => response.json())
      .then((data) => console.log(data))
      .catch((error) => {
        console.error('Error:', error);
      });
  }
  const login = () => {
    fetch('http://localhost:5000')
    .then((response) => response.json())
    .then((data) => console.log(data));
  }
  const logout = () => {
    fetch('http://localhost:5000/logout')
      .then((response) => response.json())
      .then((data) => console.log(data));
  }

  return (
    <div className="App">
      <input ref={username}></input>
      <input ref={password}></input>
      <button onClick={register}>register</button>
      <button onClick={login}>login</button>
      <button onClick={logout}>logout</button>

    </div>
  );
}

export default App;
