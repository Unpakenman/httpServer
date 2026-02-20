package mapper

import localerrors "httpServer/internal/app/errors"

func (m *mapper) ResultErrorToProto(resultError localerrors.Error) error {
	return resultError.ErrorProto()
}
