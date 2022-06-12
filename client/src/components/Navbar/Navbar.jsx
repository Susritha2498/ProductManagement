import React,{useState} from 'react'
import'./Navbar.css'

const Navbar = () => {
  const [user,setUser] = useState("")
  const [show,setShow] = useState(false)
  const [logout,setLogout] = useState(false)
  if (localStorage.key(0)!==""){
    setUser(localStorage.key(0))
  }
  const handleLogout =()=>{
    let username = localStorage.key(0)
    localStorage.removeItem(username)
    setLogout(false)
  }
  // const [home,setHome] = useState(false)
  // const [prod,setProd] = useState(false)
  // setUser(localStorage.key(0))
  // const handleHome = ()=>{
  //   if( localStorage.key(0)!=""){
  //     setHome(true)
  //   }
  //   else setHome(false)
  // }

  // const handleProducts =()=>{
  //   if( localStorage.key(0)!=""){
  //     setProd(true)
  //   }
  //   else setProd(false)
  // }
 
  return (
    <div className='app-navbar'>
      <div className='app-navlinks'>
        <h1>Producto Petrificus</h1>
        <a href={user?'/home':'/'}><h2>Home</h2></a>
        <a href={user?'/products':'/'}><h2>Products</h2></a>
      </div>
      <h2 className='app-user' onClick={setShow(true)}>{user} &nbsp; &nbsp; &nbsp; </h2>
      <a href={logout?"/":"#"}><button style={show?{opacity:"1"}:{display:"none"}} onClick={handleLogout}>logout</button></a>
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