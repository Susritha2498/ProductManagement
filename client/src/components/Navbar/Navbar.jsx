import React,{useState} from 'react'
import'./Navbar.css'

const Navbar = () => {
  const [login,setLogin] = useState(false)

  const handleClick =()=>{
    if (localStorage.key(0)!=="") {setLogin(true)}
    else {setLogin(false)}
  }

  return (
    <div className='app-navbar'>
      <div className='app-navlinks'>
        <h1>Producto Petrificus</h1>
        <a href={login?'/home':'/'} onClick={handleClick}><h2>Home</h2></a>
        <a href={login?'/products':'/'} onClick={handleClick}><h2>Products</h2></a>
      </div>
    </div>

  )
}

export default Navbar

/* div key={`user-${idx}`}>
<p>{user[2].Value}</p>
<p>{user[3].Value}</p>
<p>{user[4].Value}</p>
<p>{user[5].Value}</p>
<p>{user[6].Value}</p>
<p>{user[7].Value}</p>
<p>{user[8].Value}</p>
</div> */