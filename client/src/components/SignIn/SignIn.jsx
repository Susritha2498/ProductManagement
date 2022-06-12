import React from 'react'
import './SignIn.css'
import { Link, useNavigate } from 'react-router-dom'
const SignIn = () => {
  const gotohome = useNavigate()
  const handleLogin = async(e)=>{
    e.preventDefault()
    let email=e.target.elements.email.value
    let password=e.target.elements.password.value
    const resp=await fetch("http://localhost:5000/login",{
      method:'POST',
      headers:{
        'content-Type':'application/json'
      },
      body:JSON.stringify({
        email,password
      })
    })
    const response = await resp.json()
    console.log(response)
    if(response.Code===200){
      alert(response.Message)
      localStorage.setItem(response.Response.Username,response.Response.AuthToken)
      gotohome('/home')
    }
    else{
      alert(response.Message)
    }
  }  

  return (
    <form action="/home" className='app-signin' onSubmit={handleLogin}>
      <div className='sigin-left'>
        <h2>Producto Petrificus</h2>
        <p>Don't have an account? Register here</p>
        <Link to='/register'><button>Register</button></Link>
      </div>

      <div className='signin-form'>

        <div className='signin-input'>
          <label htmlFor="email">Email</label>
          <input type="email" id='email' />
        </div>

        <div className='signin-input'>
          <label htmlFor="password">Password</label>
          <input type="text" id='password'/>
        </div>

        <button type='submit'>Sign In</button>

      </div>
    </form>
  )
}

export default SignIn