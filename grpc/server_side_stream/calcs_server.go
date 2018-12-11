package server_side_stream

import "errors"

type ClacsServer struct {
}

func (c *ClacsServer) Calcs(req *CalReq, cs Calc_CalcsServer) error {
	if req == nil {
		return errors.New("Calcs req is nil")
	}
	x := req.GetValx()
	y := req.GetValy()
	for i := 0; i < 4; i++ {
		result := x + y + int32(i)
		res := &CalRes{Result: &result}
		if err := cs.Send(res); err != nil {
			return err
		}
	}
	return nil
}
