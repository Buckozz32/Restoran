// Получение информации о складе
func (s *Server) GetWarehouse(ctx context.Context, req *pb.GetWarehouseRequest) (*pb.GetWarehouseResponse, error) {
    // Получение информации о складе из базы данных
    warehouse, err := s.getWarehouseFromDB()
    if err != nil {
        return nil, err
    }

    // Формирование ответа клиенту
    response := &pb.GetWarehouseResponse{
        Warehouse: warehouse,
    }

    return response, nil
}

// Обновление информации о складе
func (s *Server) UpdateWarehouse(ctx context.Context, req *pb.UpdateWarehouseRequest) (*pb.UpdateWarehouseResponse, error) {
    // Получение информации о складе из базы данных
    warehouse, err := s.getWarehouseFromDB()
    if err != nil {
        return nil, err
    }

    // Обновление информации о товарах на складе
    for _, product := range req.Products {
        warehouseProduct := findWarehouseProduct(warehouse.Products, product.Id)
        if warehouseProduct != nil {
            warehouseProduct.Quantity = product.Quantity
        } else {
            warehouse.Products = append(warehouse.Products, product)
        }
    }

    // Обновление информации о складе в базе данных
    err = s.updateWarehouseInDB(warehouse)
    if err != nil {
        return nil, err
    }

    // Формирование ответа клиенту
    response := &pb.UpdateWarehouseResponse{
        Warehouse: warehouse,
    }

    return response, nil
}