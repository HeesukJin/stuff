import React, { useRef } from 'react'

/**
 * Login Page for Tradeout.
 * @return {JSX.Element} The JSX Code for home page.
 */
function HomePage() {
    const username = useRef()
    const password = useRef()

    const register = () => {
        const formData = new FormData()
        formData.append('username', username.current.value)
        formData.append('password', password.current.value)

        fetch('http://localhost:5000/register', {
            method: 'POST',
            body: formData,
        })
            .then((response) => response.json())
            .then((data) => {
                console.log(data)
                document.cookie = 'sessionID=' + data.sessionID
            })
            .catch((err) => console.log(err.message))
    }

    const login = () => {
        fetch('http://localhost:5000/')
            .then((response) => response.json())
            .then((data) => console.log(data))
    }

    const logout = () => {}

    return (
        <>
            <div>Welcome to Next.js!</div>
            <input ref={username}></input>
            <input ref={password}></input>
            <button onClick={register}>register</button>
            <button onClick={login}>login</button>
            <button onClick={logout}>logout</button>
        </>
    )
}

export default HomePage
