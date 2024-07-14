```
.
├── Dockerfile
├── Makefile                    // เก็บ script ที่จำเป็นเช่น script generate mock
├── Readme.md
├── cmd
│   ├── docs                    // swagger
│   │   └── main.go
│   └── server
│       └── main.go             // จุดร้อยรวม service กับ package ต่างใน project
├── config
│   └── config.go               // เก็บค่า config ต่างๆ ของระบบ
├── docs                        // เก็บไฟล์ที่ได้จากการ gen swagger                        
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── env
│   └── local.env.example
├── go.mod
├── go.sum
├── infrastructure
│   ├── http_client_third_party1.go
│   ├── gRPC_client_third_party2.go
│   ├── postgres.go
│   └── redis.go
├── internal
│   ├── adapter                 // จุดรวม function ที่ต่อกับ third party อื่นๆ หรือ แหล่งข้อมูลภายนอก (บางที่อาจเรียกว่า client)
│   │   ├── init.go
│   │   └── domain_adapter      // เก็บ struct request, response ของ layer adapter
│   ├── core
│   │   ├── domain              // เก็บ struct request, response ของ core domain
│   │   ├── port                // เก็บ interface
│   │   │   ├── adapter.go      // เก็บ interface ของ adapter
│   │   │   ├── mocks           // เก็บ mock function ที่ได้จากการ gen by mockery
│   │   │   │   ├── Adapter.go
│   │   │   │   ├── CacheRepository.go
│   │   │   │   ├── MiscRepository.go
│   │   │   │   └── Repository.go
│   │   │   ├── repository.go   // เก็บ interface ของ repository
│   │   │   └── service.go      // เก็บ interface ของ service
│   │   └── service
│   │       ├── init.go
│   │       ├── user_create.go
│   │       └── user_create_test.go
│   ├── handler             // plug between router and service
│   │   ├── dto             // dto (data to object) struct mapper for map data between model's struct layer and domain's struct layer
│   │   └── init.go
│   ├── repositories
│   │   ├── cache_repository
│   │   │   ├── init.go
│   │   │   └── redis_key
│   │   ├── misc_repository     // repo for wrap some function for able to mock in unit testing.
│   │   │   └── init.go
│   │   └── repository
│   │       └── init.go
│   └── router
│       └── router.go
├── pkg
│   ├── enum
│   │   └── sale_source.go
│   ├── error-response
│   │   └── error_response.go
│   ├── model
│   │   ├── v1
│   │   │   └── user.go
│   │   └── v2
│   │       └── user.go
│   └── util            // general function
│       └── parse_default_time_format.go
└── sql
    ├── 000001_init.down.sql
    └── 000001_init.up.sql
```