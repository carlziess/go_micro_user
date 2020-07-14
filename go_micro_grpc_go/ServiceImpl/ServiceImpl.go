package ServiceImpl

import (
	Services "Gomicro_grpc/service"
	"context"
	"strconv"
	"time"
)

//实现类
type ProdService struct {

}

func NewProd(id int32,name string)*Services.ProdModel{
	return &Services.ProdModel{
		ProdID: id,
		ProdName: name,
	}
}

func (this *ProdService)GetProdList(ctx context.Context,req *Services.ProdReq,rsp *Services.ProdRsp) error{
	<-time.After(time.Second*3)
	var res []*Services.ProdModel
	for i:=1;i<int(req.Size)+1;i++ {
		arg := NewProd(int32(i), "product_"+strconv.Itoa(i))
		res=append(res,arg)
	}
	rsp.Data=res
	return nil
}


func (this *ProdService) GetProdDetail(ctx context.Context, in *Services.ProdReq, out *Services.ProdDetailRsp) error {
	res :=new(Services.ProdModel)
	res.ProdName="product_"+strconv.Itoa(int(in.Pid))
	res.ProdID=in.Pid
	out.Data=res
	return nil
}