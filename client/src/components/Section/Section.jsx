import React,{useState,useEffect} from 'react'
import { useNavigate, Link } from 'react-router-dom'
import Sidebar from '../Sidebar/Sidebar'
import Product from '../Product/Product'
import AddProduct from '../AddProduct/AddProduct'
import "./Section.css"

const Section = () => {
	const [add,setAdd] = useState(false)
	const handleAdd = ()=>{
		setAdd(!add)
	}
	const gotoproducts  = useNavigate()
	const [products, setProducts] = useState([])
  
	useEffect(() => {
		async function getProducts(){
		let username = localStorage.key(0)
		let token = localStorage.getItem(username)
		const resp=await fetch('http://localhost:5000/allproducts',{
			method:'GET',
			headers:{
				'content-Type':'application/json',
				Authorization: token
			},
		})		
		const response = await resp.json()
		// console.log(response.Response)
		if(response.Code===200){
		// alert("Successfully fetched the products")
		gotoproducts('/products')
		setProducts(response.Response)
		}
		else{
		alert("Failed to fetch the products")
		} 
	}
	getProducts()

	}, [])

  return (
    <div className="app-section">
      <Sidebar/>
      <div className="app-section-products">
	  	<div className='section-add-product'>
			<h2>Products | {products.length} </h2>
			<button onClick={handleAdd}>Add product</button>
		</div>
      {products.map((item,idx)=>{
        let product = item
          return(
			<Product key={`products-${idx}`} product={item} id={product._id}/>	
          )
        })}
      </div>  
	  <div style={add?{opacity:1}:{display:"none"}}>
		<AddProduct add={add} setAdd={setAdd}/>
	  </div>
    </div>
  )
}

export default Section
