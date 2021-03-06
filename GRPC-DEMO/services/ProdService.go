package services

import (
	"context"
)

type ProdService struct {
}

func (this *ProdService) NewOrder(ctx context.Context, request *OrderRequest) (*OrderResponse, error) {
	panic("implement me")
}

//服务具体实现
func (this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	switch req.ProdArea {
	case ProdAreas_B:
		stock = 31 //B区库存31个
	case ProdAreas_C:
		stock = 50 //C区库存50
	default:
		stock = 102 //默认A区，库存102
	}
	return &ProdResponse{ProdStock: stock}, nil
}

func (this *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	Prodres := []*ProdResponse{
		&ProdResponse{ProdStock: 12},
		&ProdResponse{ProdStock: 13},
		&ProdResponse{ProdStock: 14},
	}
	return &ProdResponseList{Prodres: Prodres,}, nil
}

func (this *ProdService) GetProdInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error) {
	return &ProdModel{ProdId: 101, ProdName: "测试商品", ProdPrice: 20.5,}, nil
}
