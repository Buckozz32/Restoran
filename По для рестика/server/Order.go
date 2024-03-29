package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/Buckozz32//restaurant"
)


func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    // Создание записи в базе данных
    order := &pb.Order{
        Id:          generateOrderId(),
        CustomerId:  req.CustomerId,
        Products:    req.Products,
        TotalPrice:  calculateTotalPrice(req.Products),
        Status:      pb.OrderStatus_CREATED,
        CreatedAt:   time.Now().Unix(),
        UpdatedAt:   time.Now().Unix(),
    }
    err := s.createOrderInDB(order)
    if err != nil {
        return nil, err
    }

    // Формирование ответа клиенту
    response := &pb.CreateOrderResponse{
        Order: order,
    }

    return response, nil
}

// Получение заказа
func (s *Server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
    
    order, err := s.getOrderFromDB(req.Id)
    if err != nil {
        return nil, err
    }

    
    response := &pb.GetOrderResponse{
        Order: order,
    }

    return response, nil
}


func (s *Server) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
    
    order, err := s.getOrderFromDB(req.Id)
    if err != nil {
        return nil, err
    }

    
    order.Status = req.Status
    order.UpdatedAt = time.Now().Unix()

   
    err = s.updateOrderInDB(order)
    if err != nil {
        return nil, err
    }

   
    response := &pb.UpdateOrderResponse{
        Order: order,
    }

    return response, nil
}


func (s *Server) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
    
    err := s.deleteOrderFromDB(req.Id)
    if err != nil {
        return nil, err
    }

    
    response := &pb.DeleteOrderResponse{}

    return response, nil
}


func (s *Server) createOrderInDB(order *pb.Order) error {
    // создание заказа в базе данных
    // запросы в  таблицу "orders"
     return nil
}


func (s *Server) getOrderFromDB(id int64) (*pb.Order, error) {
    

    return &pb.Order{}, nil
}


func (s *Server) updateOrderInDB(order *pb.Order) error {
   return nil
}


func (s *Server) deleteOrderFromDB(id int64) error {
    return nil
}