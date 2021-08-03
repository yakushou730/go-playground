package errcode

import (
	pb "playground/tag-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Status struct {
	*status.Status
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

func TogRPCError(err *Error) error {
	pbErr := &pb.Error{Code: int32(err.Code()), Message: err.Msg()}
	s, _ := status.New(ToGRPCCode(err.Code()), err.Msg()).WithDetails(pbErr)
	return s.Err()
}

func ToRPCStatus(code int, msg string) *Status {
	pbErr := &pb.Error{Code: int32(code), Message: msg}
	s, _ := status.New(ToGRPCCode(code), msg).WithDetails(pbErr)
	return &Status{s}
}

func ToGRPCCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case Fail.Code():
		statusCode = codes.Internal
	case InvalidParams.Code():
		statusCode = codes.InvalidArgument
	case Unauthorized.Code():
		statusCode = codes.Unauthenticated
	case AccessDenied.Code():
		statusCode = codes.PermissionDenied
	case DeadlineExceeded.Code():
		statusCode = codes.DeadlineExceeded
	case NotFound.Code():
		statusCode = codes.NotFound
	case LimitExceed.Code():
		statusCode = codes.ResourceExhausted
	case MethodNotAllowed.Code():
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Unknown
	}
	return statusCode
}
