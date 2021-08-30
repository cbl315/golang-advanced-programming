
ch-two-hello:
	cd ch2/hello; go run . -v

ch-two-foundation:
	cd ch2/foundation; go run . -v

ch-two-typeswitch:
	cd ch2/typeswitch; go run . -v

ch-two-functioncall:
	cd ch2/functioncall; go run . -v

ch-two-cpp:
	cd ch2/cpp; go run . -v

ch-two-lib-static:
	cd ch2/lib/static; cd number;gcc -c -o number.o number.c;ar rcs libnumber.a number.o;cd -;go run .

ch-two-lib-dynamic:
	cd ch2/lib/dynamic; cd number;gcc -shared -o libnumber.so number.c;cd -;go run .

ch-two-lib-go:
	cd ch2/lib/go; go build -buildmode=c-archive -o number.a;gcc -o a.out _test_main.c number.a;./a.out

ch-two-libso-go:
	cd ch2/lib/go; go build -buildmode=c-shared -o number.so;gcc -o a.out _test_main.c number.so;./a.out

# chapter 3

ch-three-tutorial:
	# end with an empty line
	cd ch3/tutorial; go tool compile -S pkg.go

ch-four-tutorial:
	# end with an empty line
	cd ch4/tutorial; go test .

ch-four-proto:
	cd ch4/protobuf; protoc --go-grpc_out=. hello.proto