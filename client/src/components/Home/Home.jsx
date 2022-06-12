import React from 'react'
import "./Home.css"
import { useNavigate } from 'react-router-dom'

const Home = () => {
    const gotoProducts = useNavigate()
    const handleClick =()=>{
        gotoProducts('/products')
    }
  return (
    <div className="app-home">
        <div className='app-new-dresses'>
            <div className='app-dress-tag'>
                <h2>New Arrivals</h2>
                <h2>20% OFF for First Time Users</h2>
                <button onClick={handleClick}>Shop Now</button>
            </div>
        </div>

        <div className='app-new-denims'>
            <div className='app-denim-tag'>
                <h2>End of Season Sale </h2>
                <h2>FLAT 50% OFF</h2>
                <button onClick={handleClick}>Shop Now</button>
            </div>
        </div>
        <div className='app-new-dresses2'>
            <div className='app-dress2-tag'>
                <h2>New Arrivals</h2>
                <h2>Party Party Party!!!</h2>
                <button onClick={handleClick}>Shop Now</button>
            </div>
        </div>
        <div className='app-new-sarees'>
            <div className='app-saree-tag'>
                <h2> New Arrivals!!!</h2>
                <h2>Festive Season is here</h2>
                <button onClick={handleClick}>Shop Now</button>
            </div>
        </div>
        
    </div>
  )
}

export default Home