

```
project_name
├── .env
├── README.md
├── configuration
│   └── main.go
├── database
│   └── main.go
├── dockerfile
├── go.mod
├── go.sum
├── middleware
│   ├── access_control.go
├── main.go
├── post
│   ├── main.go
│   ├── main_test.go
│   ├── route.go
│   ├── struct.go
│   └── utils.go
├── comment
│   ├── main.go
│   ├── main_test.go
│   ├── route.go
│   ├── struct.go
│   └── utils.go
├── resource.go
├── route.go
└── utils
    ├── main.go
    └── struct.go
```
ด้านบนจะเป็นโครงสร้างที่ใช้ในหลายๆโปรเจค โครงสร้างนี้ design มาแบบง่ายๆ โดยจะแบ่งเป็น module ตาม folder เช่น post, comment, configuration, database และ middleware

**middleware** - จะเอาไว้จัดการพวก token หรือการเข้าถึง api

**configuration** - โหลด config จาก env ที่ระบุ

**database** - ต่อ database

**post** - service หลักของโปรเจค ด้านในจะมี controller อยู่ใน route.go และ business logic จะอยู่ใน main.go 


```
order_service
├── bootstrap
│   ├── middlewares.go
│   └── routes.go
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── grpc
│   │   ├── grpc.go
│   │   ├── order
│   │   │   └── server.go
│   │   └── protobuf
│   │       ├── orders.pb.go
│   │       └── orders_grpc.pb.go
│   ├── modules
│   │   ├── azservicebus
│   │   │   ├── module.go
│   │   │   └── service.go
│   │   ├── module.go
│   │   └── payment
│   │       ├── controller.go
│   │       ├── module.go
│   │       └── service.go
│   ├── repository
│   │   ├── models
│   │   │   ├── orders.go
│   │   │   └── reservation.model.go
│   │   └── repository.go
│   └── types
│       ├── base.response.go
│       ├── exmaple.request.go
│       └── route.go
├── main.go
├── pkg
│   ├── http
│   │   └── hqrental.http.go
│   └── utils
│       └── sanitize-html.util.go
├── protobuf
│   └── orders.proto
└── readme.md
```
ส่วน structure อันนี้เป็นโครงที่ใช้ให้เป็นแนวเดียวกับ nestjs โดยจะแยกดังนี้

**config** - เป็น service ที่ใช้ในการอ่าน configuration จาก env ที่กำหนด

**grpc** - เก็บ generated โค้ดของ protobuf และ server ที่จะให้บริการ

**modules** - เก็บ module ต่างๆที่ใช้ใน project

**repository** - เก็บ model struct และต่อ database

**types** - เก็บ constant และ struct ที่ common use