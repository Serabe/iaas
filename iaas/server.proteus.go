package iaas

import (
	"golang.org/x/net/context"
)

type iaasServiceServer struct {
}

func NewIaasServiceServer() *iaasServiceServer {
	return &iaasServiceServer{}
}
func (s *iaasServiceServer) Inc(ctx context.Context, in *IncRequest) (result *IncResponse, err error) {
	result = new(IncResponse)
	result.Result1 = Inc(in.Arg1)
	return
}
