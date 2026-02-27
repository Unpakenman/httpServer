package mapper

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	localerrors "httpServer/internal/app/errors"
)

func (m *mapper) ResultErrorToProto(code codes.Code, resultError localerrors.Error) error {
	st := status.New(code, resultError.Error())
	return st.Err()
}
