import React,{useState} from 'react'
import './Sidebar.css'
import {HiOutlineFilter,HiOutlineHome} from 'react-icons/hi'
import {AiOutlineGift,AiOutlinePlusSquare,AiOutlineRightCircle} from 'react-icons/ai'
import {BsX,BsList} from 'react-icons/bs'
const Sidebar = () => {
  const [logout,setLogout] = useState(false)
  const handleLogout = ()=>{
    let username = localStorage.key(0)
    localStorage.removeItem(username)
    setLogout(true)
  }
  return (
    <div className='app-products-sidebar'>
        <BsX size={30} className='side-close'/>
        <BsList size={30} className='side-icons' />
        <a href={logout?"/":"#"}><HiOutlineHome size={30} className='side-icons' onClick={handleLogout}/></a>
        <HiOutlineFilter size={30} className='side-icons'/>
        <AiOutlineGift size={30} className='side-icons' />
        <AiOutlinePlusSquare size={30} className='side-icons'/>
        <AiOutlineRightCircle size={30} className='side-icons'/>
    </div>
  )
}

export default Sidebar