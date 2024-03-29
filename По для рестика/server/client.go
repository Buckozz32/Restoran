package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
    "github.com/golang/protobuf/ptypes/empty"
    pb "github.com/Buckozz32//restaurant"
)

func createOrder(client pb.RestaurantClient, productOrder []*pb.ProductOrder, comment string) (*pb.CreateOrderResponse, error) {
    return client.CreateOrder(context.Background(), &pb.CreateOrderRequest{ProductOrder: productOrder, Comment: comment})
}

func getOrder(client pb.RestaurantClient, id int32) (*pb.GetOrderResponse, error) {
    return client.GetOrder(context.Background(), &pb.GetOrderRequest{Id: id})
}

func updateOrder(client pb.RestaurantClient, id int32, status string, comment string) (*pb.UpdateOrderResponse, error) {
    return client.UpdateOrder(context.Background(), &pb.UpdateOrderRequest{Id: id, Status: status, Comment: comment})
}

func deleteOrder(client pb.RestaurantClient, id int32) (*empty.Empty, error) {
    return client.DeleteOrder(context.Background(), &pb.DeleteOrderRequest{Id: id})
}

func getProducts(client pb.RestaurantClient) (*pb.GetProductsResponse, error) {
    return client.GetProducts(context.Background(), &empty.Empty{})
}

func getProduct(client pb.RestaurantClient, id int32) (*pb.GetProductResponse, error) {
    return client.GetProduct(context.Background(), &pb.GetProductRequest{Id: id})
}

func updateProduct(client pb.RestaurantClient, id int32, name string, description string, price float32, category string, image string) (*pb.UpdateProductResponse, error) {
    return client.UpdateProduct(context.Background(), &pb.UpdateProductRequest{Id: id, Name: name, Description: description, Price: price, Category: category, Image: image})
}

func deleteProduct(client pb.RestaurantClient, id int32) (*empty.Empty, error) {
    return client.DeleteProduct(context.Background(), &pb.DeleteProductRequest{Id: id})
}

func getWarehouse(client pb.RestaurantClient) (*pb.GetWarehouseResponse, error) {
    return client.GetWarehouse(context.Background(), &empty.Empty{})
}

func updateWarehouse(client pb.RestaurantClient, id int32, productId int32, quantity int32, reserved int32) (*pb.UpdateWarehouseResponse, error) {
    return client.UpdateWarehouse(context.Background(), &pb.UpdateWarehouseRequest{Id: id, ProductId: productId, Quantity: quantity, Reserved: reserved})
}

func getSuppliers(client pb.RestaurantClient) (*pb.GetSuppliersResponse, error) {
    return client.GetSuppliers(context.Background(), &empty.Empty{})
}

func getSupplier(client pb.RestaurantClient, id int32) (*pb.GetSupplierResponse, error) {
    return client.GetSupplier(context.Background(), &pb.GetSupplierRequest{Id: id})
}

func createSupplier(client pb.RestaurantClient, name string, address string, phone string, email string) (*pb.CreateSupplierResponse, error) {
    return client.CreateSupplier(context.Background(), &pb.CreateSupplierRequest{Name: name, Address: address, Phone: phone, Email: email})
}

func updateSupplier(client pb.RestaurantClient, id int32, name string, address string, phone string, email string) (*pb.UpdateSupplierResponse, error) {
    return client.UpdateSupplier(context.Background(), &pb.UpdateSupplierRequest{Id: id, Name: name, Address: address, Phone: phone, Email: email})
}

func deleteSupplier(client pb.RestaurantClient, id int32) (*empty.Empty, error) {
    return client.DeleteSupplier(context.Background(), &pb.DeleteSupplierRequest{Id: id})
}

func getDeliveries(client pb.RestaurantClient) (*pb.GetDeliveriesResponse, error) {
    return client.GetDeliveries(context.Background(), &empty.Empty{})
}

func createDelivery(client pb.RestaurantClient, supplierId int32, productOrder []*pb.ProductOrder) (*pb.CreateDeliveryResponse, error) {
    return client.CreateDelivery(context.Background(), &pb.CreateDeliveryRequest{SupplierId: supplierId, ProductOrder: productOrder})
}

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewRestaurantClient(conn)

    productOrder := []*pb.ProductOrder{
        {ProductId: 1, Quantity: 2},
        {ProductId: 2, Quantity: 3},
    }
    response, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{ProductOrder: productOrder, Comment: "Test order"})
    if err != nil {
        log.Fatalf("could not create order: %v", err)
    }
    log.Printf("Order created: %v", response)
}