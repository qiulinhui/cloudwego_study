.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl/ --type RPC --module cloudwego_study/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server -I ../../idl/ --type RPC --module cloudwego_study/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift