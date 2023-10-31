go_protoc:
	protoc --go_out=./Gateway/services --go-grpc_out=./Gateway/services  ./Gateway/proto/owner.proto
	protoc --go_out=./Gateway/services --go-grpc_out=./Gateway/services  ./Gateway/proto/veterinary.proto
	protoc --go_out=./Gateway/services --go-grpc_out=./Gateway/services  ./Gateway/proto/test.proto

init: go_protoc
	linkerd install --crds | kubectl apply -f -
	linkerd install --set proxyInit.runAsRoot=true | kubectl apply -f -
	linkerd check

run: 
	docker compose pull
	linkerd inject k8s.yaml | kubectl apply -f -

stop: 
	linkerd uninject k8s.yaml | kubectl delete -f -

tests: 
	dotnet test OwnerService

uninstall:
	linkerd uninstall | kubectl delete -f -
	
push:
	docker compose build
	docker compose push