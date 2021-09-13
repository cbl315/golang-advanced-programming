
ch2-hello:
	cd ch2/hello; go run . -v

ch2-foundation:
	cd ch2/foundation; go run . -v

ch2-typeswitch:
	cd ch2/typeswitch; go run . -v

ch2-functioncall:
	cd ch2/functioncall; go run . -v

ch2-cpp:
	cd ch2/cpp; go run . -v

ch2-lib-static:
	cd ch2/lib/static; cd number;gcc -c -o number.o number.c;ar rcs libnumber.a number.o;cd -;go run .

ch2-lib-dynamic:
	cd ch2/lib/dynamic; cd number;gcc -shared -o libnumber.so number.c;cd -;go run .

ch2-lib-go:
	cd ch2/lib/go; go build -buildmode=c-archive -o number.a;gcc -o a.out _test_main.c number.a;./a.out

ch2-libso-go:
	cd ch2/lib/go; go build -buildmode=c-shared -o number.so;gcc -o a.out _test_main.c number.so;./a.out

# chapter 3

ch3-tutorial:
	# end with an empty line
	cd ch3/tutorial; go tool compile -S pkg.go

ch4-tutorial:
	# end with an empty line
	cd ch4/tutorial; go test .

ch4-proto:
	cd ch4/protobuf; protoc --go-grpc_out=. hello.proto

ch4-grpc:
	cd ch4/gRPC; protoc --go-grpc_out=. hello.proto --go_out=./;

ch4-pubsub:
	cd ch4/pubsub; protoc --go-grpc_out=. pubsub.proto --go_out=./;

ch5-intro:
	cd ch5/intro; go run hello.go

ch5-rate-limit:
	cd ch5/rateLimit; go run rateLimit.go