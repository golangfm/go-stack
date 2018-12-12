package both_stream

import (
	"io"
	"log"
	"strconv"
)

type ChatService struct {
}

func (cs *ChatService) BidStream(cb Chat_BidStreamServer) error {
	ctx := cb.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("recv the context done.")
			return ctx.Err()
		default:
			rawData, err := cb.Recv()
			if err == io.EOF {
				log.Println("client send eof.")
				return nil
			}
			if err != nil {
				log.Println("recv error:" + err.Error())
				return err
			}
			switch rawData.GetInput() {
			case "done\n":
				log.Println("recv done.")
				output := "recv done."
				if err := cb.Send(&Res{Output: &output}); err != nil {
					return err
				}
				return nil
			case "return\n":
				log.Println("recv return.")
				for i := 0; i < 10; i++ {
					output := "data flow with index" + strconv.Itoa(i)
					if err := cb.Send(&Res{Output: &output}); err != nil {
						return err
					}
				}
				return nil
			default:
				log.Println("recv data:" + rawData.GetInput())
				output := "server return:" + rawData.GetInput()
				if err := cb.Send(&Res{Output: &output}); err != nil {
					return err
				}
			}
		}
	}
}
