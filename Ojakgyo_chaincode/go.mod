module github.com/nu-jey/blocker_chaincode
//module - version control을 위한 url을 저장한다.
//이 프로젝트가 github.com/nu-jey/blocker_chaincode 라는 이름의 Go 모듈을 가짐을 선언
//모듈은 코드의 단위로써, 일종의 패키지 모음이라고 생각할 수 있음

go 1.19
// go 1.14 - golang ver.

require github.com/hyperledger/fabric-contract-api-go v1.2.0
//require - 프로젝트에서 사용하는 모든 의존 모듈 리스트
//github.com/hyperledger/fabric-contract-api-go 모듈의 버전 v1.2.0을 이 프로젝트가 사용한다고 선언, 이 프로젝트는 해당 버전의 모듈에 의존

//프로젝트가 사용하는 다른 모듈들과 해당 모듈들의 버전을 나열
//indirect - 명시적으로 사용하지는 않고, 내부 모듈에서 사용하는 의존성, 해당 의존성이 직접 사용되지 않고, 내부 의존성에서 사용된다는 것을 나타냄
require (
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/gobuffalo/envy v1.10.1 // indirect
	github.com/gobuffalo/packd v1.0.1 // indirect
	github.com/gobuffalo/packr v1.30.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20220720122508-9207360bbddd // indirect
	github.com/hyperledger/fabric-protos-go v0.0.0-20220613214546-bf864f01d75e // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	golang.org/x/net v0.0.0-20220708220712-1185a9018129 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220719170305-83ca9fad585f // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace - git version이 아닌 local version of dependency 지칭
//원격 버전 대신 로컬에서 개발 중인 버전의 의존성을 사용할 때 지정됨, 이 프로젝트는 원격 버전이 아닌 로컬 버전의 의존성을 사용