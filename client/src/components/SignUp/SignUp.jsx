import React from 'react'
import './SignUp.css'
import { useNavigate } from 'react-router-dom'
const SignUp = () => {
  const gotologin = useNavigate()
  const handleRegister = async(e)=>{
      e.preventDefault()
      let username=e.target.elements.name.value
      let email=e.target.elements.email.value
      let password=e.target.elements.password.value
      let phone=e.target.elements.phone.value
      let city=e.target.elements.city.value
      const resp=await fetch("http://localhost:5000/register",{
        method:'POST',
        headers:{
          'content-Type':'application/json'
        },
        body:JSON.stringify({
          username,email,password,phone,city
        })
      })
      const response = await resp.json()
      console.log(response);
      if(response.Code===200){
        alert(response.Message)
        gotologin('/')
      }
      else{
        alert(response.Message)
      }
  
  }
  return (
    <form action="/" className='app-register' onSubmit={handleRegister}>
      <div className='register-left'>
        <h2>Producto Petrificus</h2>
        <p>Wide range of offers this season!!Hurry up?? Limited Stock!!!</p>
        <button>Shop now</button>
      </div>

      <div className='register-form'>

        <div className='register-input'>
          <label htmlFor="name">Username</label>
          <input type="text" id='name' />
        </div>

        <div className='register-input'>
          <label htmlFor="email">Email</label>
          <input type="email" id='email'/>
        </div>

        <div className='register-input'>
          <label htmlFor="password">Password</label>
          <input type="text" id='password'/>
        </div>

        <div className='register-input'>
          <label htmlFor="phone">Phone</label>
          <input type="number" id='phone'/>
        </div>

        <div className='register-input'>
          <label htmlFor="city">City</label>
          <input type="text" id='city'/>
        </div>

        <button type='submit'>Register</button>
      </div>

  </form>
  )
}

export default SignUp
