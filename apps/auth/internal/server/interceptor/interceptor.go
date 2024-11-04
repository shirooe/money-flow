package interceptor

import "google.golang.org/grpc"

func New(
	auth authInterceptor,
) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		auth.UnaryInterceptor(),
	}
}
