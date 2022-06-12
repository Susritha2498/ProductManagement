import React,{useState} from 'react'
import EditProduct from '../EditProduct/EditProduct'
import {RiDeleteBin5Line,RiEditBoxLine} from 'react-icons/ri'
import { useNavigate } from 'react-router-dom'

const Product = ({product,id}) => {
  const gotoProducts = useNavigate()
    const [edit,setEdit] = useState(false)
    const handleEdit=()=>{
      setEdit(true)
    }
    const handleDelete = async(e)=>{
      let username = localStorage.key(0)
      let token = localStorage.getItem(username)
      console.log(token);
      const resp=await fetch(`http://localhost:5000/delete/product/${id}`,{
      method:'DELETE',
      headers:{
        'content-Type':'application/json',
        Authorization: token,
      },
    })
    const response = await resp.json()
    console.log(response);
    if(response.Code===200){
      alert(response.Message)
      gotoProducts('/products')
    }
    else{
      alert(response.Message)
    } 
  }
  return (
    <>
    <div className='products-image' >
        <img src={product.productImage} alt="sectionImage"/>
        <h3>{product.title}</h3>
        <p>Rs. {product.price} </p>
        <span className='span1'>{Number(Math.round((product.rating)+'e2')+'e-2')}✩ &nbsp; &nbsp; &nbsp; {product.totalRatings} ratings</span>
        <span className='span2'>
        <RiEditBoxLine size={20} className='edit-icon' style={{marginRight:"10px"}} onClick={handleEdit}/>
        <RiDeleteBin5Line size={20} className='delete-icon' onClick={handleDelete}/>
        </span>
    </div>
      <div style={edit?{opacity:'1'}:{display:'none'}} className="edit-form">
        <EditProduct id={product.id} product={product} edit={edit} setEdit={setEdit}/>
      </div> 
    </>
  )
}

export default Product