// Получение списка поставщиков
func (s *Server) GetSuppliers(ctx context.Context, req *pb.GetSuppliersRequest) (*pb.GetSuppliersResponse, error) {
	// Получение списка поставщиков из базы данных
	suppliers, err := s.getSuppliersFromDB()
	if err != nil {
		return nil, err
	}

	// Формирование ответа клиенту
	response := &pb.GetSuppliersResponse{
		Suppliers: suppliers,
	}

	return response, nil
}

// Получение информации о поставщике
func (s *Server) GetSupplier(ctx context.Context, req *pb.GetSupplierRequest) (*pb.GetSupplierResponse, error) {
	// Получение информации о поставщике из базы данных
	supplier, err := s.getSupplierFromDB(req.Id)
	if err != nil {
		return nil, err
	}

	// Формирование ответа клиенту
	response := &pb.GetSupplierResponse{
		Supplier: supplier,
	}

	return response, nil
}

// Создание нового поставщика
func (s *Server) CreateSupplier(ctx context.Context, req *pb.CreateSupplierRequest) (*pb.CreateSupplierResponse, error) {
	// Создание нового поставщика в базе данных
	supplier := &pb.Supplier{
		Id:          generateSupplierId(),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
	err := s.createSupplierInDB(supplier)
	if err != nil {
		return nil, err
	}

	// Формирование ответа клиенту
	response := &pb.CreateSupplierResponse{
		Supplier: supplier,
	}

	return response, nil
}

// Обновление информации о поставщике
func (s *Server) UpdateSupplier(ctx context.Context, req *pb.UpdateSupplierRequest) (*pb.UpdateSupplierResponse, error) {
	// Получение информации о поставщике из базы данных
	supplier, err := s.getSupplierFromDB(req.Id)
	if err != nil {
		return nil, err
	}

	// Обновление информации о поставщике
	supplier.Name = req.Name
	supplier.Description = req.Description
	supplier.UpdatedAt = time.Now().Unix()

	// Обновление информации о поставщике в базе данных
	err = s.updateSupplierInDB(supplier)
	if err != nil {
		return nil, err
	}

	// Формирование ответа клиенту
	response := &pb.UpdateSupplierResponse{
		Supplier: supplier,
	}

	return response, nil
}

// Удаление поставщика
func (s *Server) DeleteSupplier(ctx context.Context, req *pb.DeleteSupplierRequest) (*pb.DeleteSupplierResponse, error) {
	// Удаление поставщика из базы данных
	err := s.deleteSupplierFromDB(req.Id)
	if err != nil {
		return nil, err
	}

	// Формирование ответа клиенту
	response := &pb.DeleteSupplierResponse{}

	return response, nil
}

// Создание новой поставки
func (s *Server) CreateDelivery(ctx context.Context, req *pb.CreateDeliveryRequest) (*pb.CreateDeliveryResponse, error) {
	// Получение информации о поставщике из базы данных
	supplier, err := s.getSupplierFromDB(req.SupplierId)
	if err != nil {
		return nil, err
	}

	// Получение текущей информации о складе из базы данных
	warehouse, err := s.getWarehouseFromDB()
	if err != nil {
		return nil, err
	}

	// Создание новой поставки
	delivery := &pb.Delivery{
		Id:         generateDeliveryId(),
		SupplierId: req.SupplierId,
		Supplier:   supplier,
		Products:   req.Products,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	// Обновление информации о товарах на складе
	for _, product := range req.Products {
		warehouseProduct := findWarehouseProduct(warehouse.Products, product.Id)
		if warehouseProduct != nil {
			warehouseProduct.Quantity += product.Quantity
		} else {
			warehouse.Products = append(warehouse.Products, &pb.WarehouseProduct{
				Id:       product.Id,
				Quantity: product.Quantity,
			})
		}
	}

	// Сохранение новой поставки в базе данных
	err = s.createDeliveryInDB(delivery)
	if err != nil {
		return nil, err
	}

	// Обновление информации о складе в базе данных
	err = s.updateWarehouseInDB(warehouse)
	if err != nil {
		return nil, err
	}

	// Формирование ответа клиенту
	response := &pb.CreateDeliveryResponse{
		Delivery: delivery,
	}

	return response, nil
}