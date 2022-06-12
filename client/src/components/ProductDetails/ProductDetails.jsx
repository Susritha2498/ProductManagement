import React from 'react'
import './ProductDetails.css'

const ProductDetails = ({id,products}) => {
  let slug = window.location.pathname.split("/")[2]
  let reqdata = products.filter((product)=>{return product.id===id})
  console.log(reqdata)
  let product=reqdata

  return (
    <div>
        <img src={product.productImage} alt="ProductImage" style={{width:'500px',height:"500px",marginTop:"20px",borderRadius:"10px",boxShadow:"0px 0px 6px grey"}}/>
        <div>
          <h1>{product.title}</h1>
          <p>{product.tagline}</p>
          <p>{product.price}</p>
          <p>{product.rating}</p>
          <p>{product.totalRatings}</p>
          <p>{product.description}</p>
          <p>{product.size}</p>
          <p>{product.productDetails}</p> 
        </div>
      
    </div>
  )
}

export default ProductDetails
