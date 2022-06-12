import React from 'react'
import './EditProduct.css'
import { useNavigate } from 'react-router-dom'
import {BsX} from 'react-icons/bs'
const EditProduct = ({product,id,edit,setEdit}) => {
  const gotoProducts = useNavigate()
  const handleClose=()=>{
    setEdit(!edit)
  }
  const handleUpdate= async(e)=>{
      e.preventDefault()
      let username = localStorage.key(0)
      let token = localStorage.getItem(username)
      let title=e.target.elements.title.value
      let tagline=e.target.elements.tagline.value
      let rating = e.target.elements.rating.value
      let totalRatings =e.target.elements.totalratings.value
      let price = e.target.elements.price.value
      let productImage = e.target.elements.imageurl.value
      const resp=await fetch(`http://localhost:5000/update/product/${id}`,{
      method:'PUT',
      headers:{
        'content-Type':'application/json',
        Authorization: token,
      },
      body:JSON.stringify({
       id,title,tagline,rating,totalRatings,price,productImage
      })
    })
    const response = await resp.json()
    console.log(response);
    if(response.Code===200){
      alert(response.Message)
      gotoProducts('/products')
    }
    else{
      alert(response.Message)
      setEdit(!edit)
    } 
  }


  return (
    <form action="/products" className='editproduct-edit-form' onSubmit={handleUpdate}>
      <BsX size={20} className="edit-close" onClick={handleClose}/>

      <div className='edit-input'>
      <label htmlFor="title">Title</label>
      <input type="text" id='title' defaultValue={product.title}/>
      </div>

      <div className='edit-input'>
      <label htmlFor="tagline">Tagline</label>
      <input type="text" id='tagline' defaultValue={product.tagline}/>
      </div>

      <div className='edit-input'>
      <label htmlFor="rating">Rating</label>
      <input type="text" id='rating' defaultValue={product.rating}/>
      </div>

      <div className='edit-input'>
      <label htmlFor="totalratings">Total Ratings</label>
      <input type="text" id='totalratings' defaultValue={product.totalRatings}/>
      </div>

      <div className='edit-input'>
      <label htmlFor="price">Price</label>
      <input type="text" id='price' defaultValue={product.price}/>
      </div>

      <div className='edit-input'>
      <label htmlFor="imageurl">Image link</label>
      <input type="text" id='imageurl' defaultValue={product.productImage}/>
      </div>

      <button type='submit'>Edit Product </button>

    </form>
  )
}

export default EditProduct
