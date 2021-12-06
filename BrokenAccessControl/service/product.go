package service

import (
	"BrokenAuthentication/dto"
)

type ProductService interface {
	CreateProduct()
	GetProduct()
	GetProducts()
	UpdateProduct()
	DeleteProduct()

}


type ProductServiceImpl struct {

}

var Product *ProductServiceImpl


func (pS *ProductServiceImpl) CreateProduct(request dto.CreateProductRequest){

}

func (pS *ProductServiceImpl) GetProduct(id uint){

}

func (pS *ProductServiceImpl) GetProducts(){

}

func (pS *ProductServiceImpl) UpdateProduct(request dto.UpdateProductRequest){

}

func (pS *ProductServiceImpl) DeleteProduct(id uint){

}