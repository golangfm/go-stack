package grpc_simple

import (
	context "context"
	"errors"
)

type SimpleCalcServer struct {
}

func (a *SimpleCalcServer) Calc(ctx context.Context, req *CalReq) (*CalRes, error) {
	if req == nil {
		return nil, errors.New("Calc req params is nil.")
	}
	var result int32
	switch req.GetOp() {
	case "+":
		result = req.GetValx() + req.GetValy()
	case "-":
		result = req.GetValx() - req.GetValy()
	case "*":
		result = req.GetValx() * req.GetValy()
	case "/":
		result = req.GetValx() / req.GetValy()
	case "%":
		result = req.GetValx() % req.GetValy()
	default:
		return nil, errors.New("unsupport op.")
	}
	return &CalRes{Result: &result}, nil
}
