package main

import (
    "context"
    "log"
    "net"
	"database/sql"
	

    "google.golang.org/grpc"
    "github.com/golang/protobuf/ptypes/empty"
    pb "github.com/yourusername/yourproject/restaurant"
	
)

type server struct{
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
    Server := // подключение к бд
    return &Server{db: db}
}

func (s *server) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    
    return &pb.CreateOrderResponse{}, nil
}

func (s *server) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
   
    return &pb.GetOrderResponse{}, nil
}

func (s *server) UpdateOrder(ctx context.Context, in *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
    
    return &pb.UpdateOrderResponse{}, nil
}

func (s *server) DeleteOrder(ctx context.Context, in *pb.DeleteOrderRequest) (*empty.Empty, error) {
   
    return &empty.Empty{}, nil
}

func (s *server) GetProducts(ctx context.Context, in *empty.Empty) (*pb.GetProductsResponse, error) {
    
    return &pb.GetProductsResponse{}, nil
}

func (s *server) GetProduct(ctx context.Context, in *pb.GetProductRequest) (*pb.GetProductResponse, error) {
    
    return &pb.GetProductResponse{}, nil
}

func (s *server) UpdateProduct(ctx context.Context, in *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
    
    return &pb.UpdateProductResponse{}, nil
}

func (s *server) DeleteProduct(ctx context.Context, in *pb.DeleteProductRequest) (*empty.Empty, error) {
  
    return &empty.Empty{}, nil
}

func (s *server) GetWarehouse(ctx context.Context, in *empty.Empty) (*pb.GetWarehouseResponse, error) {
    
    return &pb.GetWarehouseResponse{}, nil
}

func (s *server) UpdateWarehouse(ctx context.Context, in *pb.UpdateWarehouseRequest) (*pb.UpdateWarehouseResponse, error) {
    
    return &pb.UpdateWarehouseResponse{}, nil
}

func (s *server) GetSuppliers(ctx context.Context, in *empty.Empty) (*pb.GetSuppliersResponse, error) {
    
    return &pb.GetSuppliersResponse{}, nil
}

func (s *server) GetSupplier(ctx context.Context, in *pb.GetSupplierRequest) (*pb.GetSupplierResponse, error) {
   
    return &pb.GetSupplierResponse{}, nil
}

func (s *server) CreateSupplier(ctx context.Context, in *pb.CreateSupplierRequest) (*pb.CreateSupplierResponse, error) {
   
    return &pb.CreateSupplierResponse{}, nil
}

func (s *server) UpdateSupplier(ctx context.Context, in *pb.UpdateSupplierRequest) (*pb.UpdateSupplierResponse, error) {
    
    return &pb.UpdateSupplierResponse{}, nil
}

func (s *server) DeleteSupplier(ctx context.Context, in *pb.DeleteSupplierRequest) (*empty.Empty, error) {
    
    return &empty.Empty{}, nil
}

func (s *server) GetDeliveries(ctx context.Context, in *empty.Empty) (*pb.GetDeliveriesResponse, error) {
   
    return &pb.GetDeliveriesResponse{}, nil
}

func (s *server) CreateDelivery(ctx context.Context, in *pb.CreateDeliveryRequest) (*pb.CreateDeliveryResponse, error) {
   
    return &pb.CreateDeliveryResponse{}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterRestaurantServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}