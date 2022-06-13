import React from 'react'
import './AddProduct.css'
import { useNavigate } from 'react-router-dom'
import {BsX} from 'react-icons/bs'
const AddProduct = ({add,setAdd,getProducts}) => {
  const gohome = useNavigate()
  const handleClose=()=>{
    setAdd(!add)
  }
    const handleAdd= async(e)=>{
      e.preventDefault()
      let username = localStorage.key(0)
      let token = localStorage.getItem(username)
      let title=e.target.elements.title.value
      let tagline=e.target.elements.tagline.value
      let rating = e.target.elements.rating.value
      let totalRatings =e.target.elements.totalratings.value
      let price = e.target.elements.price.value
      let productImage = e.target.elements.imageurl.value
    const resp=await fetch("http://localhost:5000/add/product",{
      method:'POST',
      headers:{
        'content-Type':'application/json',
        Authorization: token,
      },
      body:JSON.stringify({
        title,tagline,rating,totalRatings,price,productImage
      })
    })
    const response = await resp.json()
    if(response.Code===200){
      alert(response.Message)
      setAdd(!add)
      getProducts()
    }else if(response.Code===404){
      alert(response.Message)
      gohome('/')
    }
    else{
      alert(response.Message)
      setAdd(!add)
    } 
  }
    return (
      <form action="/products" className='addproduct-add-form' onSubmit={handleAdd}>
        <BsX size={20} className="add-close" onClick={handleClose}/>
  
        <div className='add-input'>
        <label htmlFor="title">Title</label>
        <input type="text" id='title'/>
        </div>
  
        <div className='add-input'>
        <label htmlFor="tagline">Tagline</label>
        <input type="text" id='tagline'/>
        </div>
  
        <div className='add-input'>
        <label htmlFor="rating">Rating</label>
        <input type="text" id='rating'/>
        </div>
  
        <div className='add-input'>
        <label htmlFor="totalratings">Total Ratings</label>
        <input type="text" id='totalratings'/>
        </div>
  
        <div className='add-input'>
        <label htmlFor="price">Price</label>
        <input type="text" id='price'/>
        </div>
  
        <div className='add-input'>
        <label htmlFor="imageurl">Image link</label>
        <input type="text" id='imageurl'/>
        </div>
  
        <button type='submit'>Add Product</button>
      </form>
  )
}

export default AddProduct
