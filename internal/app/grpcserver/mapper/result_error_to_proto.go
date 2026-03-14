package mapper

import (
	"errors"
	pbcommon "github.com/Unpakenman/protos/gen/go/sso/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/protoadapt"
	appErrors "httpServer/internal/app/errors"
	localerrors "httpServer/internal/app/errors"
)

func (m *mapper) toProtoError(code codes.Code, errorMessage string, details proto.Message) error {
	st := status.New(code, errorMessage)
	ds, err := st.WithDetails(protoadapt.MessageV1Of(details))
	if err != nil {
		return st.Err()
	}

	return ds.Err()
}

func (m *mapper) ResultErrorToProtoError(resultError localerrors.Error) error {
	errorMessage := resultError.Error()

	if errors.Is(resultError, appErrors.ErrInvalidDateTime) {
		details := &pbcommon.DateTimeErrorDetails{}
		return m.toProtoError(codes.InvalidArgument, errorMessage, details)
	}
	return resultError.ErrorProto()
}
