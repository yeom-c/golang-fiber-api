# Golnag Fiber API

### 소개

> Go Fiber 프레임워크를 사용하여 백엔드 API 를 제공합니다.

- [Go](https://go.dev)
- [Gin](https://gin-gonic.com/)
- [Air: Live Reload Module](https://github.com/cosmtrek/air)

### 폴더 구조

```
.
├── config          # 설정 관련 데이터들을 제공
│   ├── constant    # constant 정의
│   ├── enum        # enum 정의
│   └── env         # 환경 변수 정의
├── controller      # route 요청 핸들 컨트롤러
├── database        # db connection struct
├── middleware      # request/response 전,후 처리 미들웨어
├── model
│   ├── dto         # 전달용 데이터 struct
│   └── entity      # 데이터베이스 entity struct
├── repository      # 데이터베이스 CRUD 로직 처리
├── route           # api route
├── service         # 비지니스 로직 처리
├── tool            # 개발용 tool
│   └── cmd         # cmd tool
└── util            # util 함수 모음
```

### 요청 응답 처리
route <=> middleware <=> controller <=> service <=> repository <=> database
