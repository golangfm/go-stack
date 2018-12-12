package client_side_stream

import (
	"io"
)

type SumService struct {
}

func (s *SumService) Sum(cs Calc_SumServer) error {
	var sum int32
	for {
		rawData, err := cs.Recv()
		if err == io.EOF {
			cs.SendAndClose(&Res{Result: &sum})
		}
		if err != nil {
			return err
		}
		sum += rawData.GetVal()
	}
}
