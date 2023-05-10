# 연톡 <YTTP: Yonsei Talk Transfer Protocol>

TCP 기반으로 L7 프로토콜입니다. 

2명의 사용자가 메시지를 주고받는 상황을 가정합니다. 

## 실행 방법 

1. 서버 실행 
```
    go run server/server.go
```
2. 클라이언트 실행
```
    go run client.go
```

## 작업 내역 
- [x] 사용자 아이디 입력 받기 
- [x] Rabin FingerPrint 로직 
- [x] Content-lenght, date-time header 구현 
- [x] 페이로드 입력 받음  
- [x] 바이너리 코드 형태로 전송 
- [ ] 사용자 아이디 유효성 체크 
- [ ] ping / pong 구현  
- [ ] 한글 영어 language 구분 
- [ ] 다른 이용에게도 페이로드 전송    
