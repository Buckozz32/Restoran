package main

import (
	""
	"context"
	"time"
)

// Получение списка продуктов
func (s *Server) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
    // Получение списка продуктов из базы данных
    products, err := s.getProductsFromDB()
    if err != nil {
        return nil, err
    }

    // Формирование ответа клиенту
    response := &pb.GetProductsResponse{
        Products: products,
    }

    return response, nil
}

// Получение продукта
func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
    // Получение продукта из базы данных
    product, err := s.getProductFromDB(req.Id)
    if err != nil {
        return nil, err
    }

    // Формирование ответа клиенту
    response := &pb.GetProductResponse{
        Product: product,
    }

    return response, nil
}

// Обновление продукта
func (s *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
    // Получение продукта из базы данных
    product, err := s.getProductFromDB(req.Id)
    if err != nil {
        return nil, err
    }

    // Обновление информации о продукте
    product.Name = req.Name
    product.Description = req.Description
    product.Price = req.Price
    product.Category = req.Category
    product.Image = req.Image
    product.UpdatedAt = time.Now().Unix()

    // Обновление продукта в базе данных
    err = s.updateProductInDB(product)
    if err != nil {
        return nil, err
    }

    // Формирование ответа клиенту
    response := &pb.UpdateProductResponse{
        Product: product,
    }

    return response, nil
}