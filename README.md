----------------------------------------------
## ibk
----------------------------------------------
- 유용한 단축키 : shift + ctrl + p1(메뉴보기)
- 폰트지정 : #의 갯수로 표시
  # 1개
  ## 2개
  ### 3개
  #### 4개
- server : origin(push), local : master(commit)

----------------------------------------------
## golang의 특징
----------------------------------------------
- := 변수선언+초기화
- 함수 리턴의 변수가 여러개 가능, 리턴한 값은 무조건 사용해야 함
- marshal(json -> object), unmarshal(object -> json)
- key : string, value : byte-array

- 실행 : > go run -f
- 빌드 : > go build -f

- local : master : 저장 : git commit (all)
- server : origin : 업로드 : git push

----------------------------------------------
## chaincode의 종류
----------------------------------------------
•CSCC : Configuration System Chaincode

•LSCC : Life Cycle System Chaincode

•QSCC : Query System Chaincode

•ESCC : Endorser System Chaincode 

•VSCC : Validator System Chaincode



----------------------------------------------
## package shim 
----------------------------------------------
◦func (stub *ChaincodeStub) GetState(key string) ([]byte, error)
◦func (stub *ChaincodeStub) PutState(key string, value []byte) error
◦func (stub *ChaincodeStub) DelState(key string) error

◦func (stub *ChaincodeStub) GetHistoryForKey(key string) (HistoryQueryIteratorInterface, error)

----------------------------------------------
## chaincode의 형태
----------------------------------------------
package main // go 파일이 위치한 패키지를 나타냄

import () // import하려는 패키지를 나타냄

type SmartContract struct {} // struct를 나타냄

func (t *SmartContract)Invoke() {} // 중요! StateDB에 Update/Query하기 위해 반드시 호출

func (t *SmartContract)Init() {} Instantiate/Upgrade 수행시 호출

func main() {} //golang의 시작점(java의 main함수와 비슷하다고 보면 됨)

