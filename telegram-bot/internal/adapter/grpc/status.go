package grpc

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func AsReason(err error, reason string) (*errdetails.ErrorInfo, bool) {
	st := status.Convert(err)

	for _, detail := range st.Details() {
		switch t := detail.(type) {
		case *errdetails.ErrorInfo:
			if t.Reason == reason {
				return t, true
			}
		}
	}

	return nil, false
}
