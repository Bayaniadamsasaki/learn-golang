# 📚 Golang Learning Path - Dari Beginner hingga Expert

Repositori ini berisi struktur pembelajaran Golang yang lengkap dan terorganisir dari level beginner hingga expert, dilengkapi dengan project examples untuk praktik langsung.

## 🏗️ Struktur Utama

```
learn-golang/
├── 01-beginner/          # Level Pemula (8 contoh)
├── 02-intermediate/      # Level Menengah (6 contoh) 
├── 03-advanced/         # Level Lanjutan (5 contoh)
├── 04-expert/           # Level Expert (4 contoh)
└── examples/            # Project Examples (8 project)
    ├── 01-beginner/     # 2 project beginner
    ├── 02-intermediate/ # 2 project intermediate
    ├── 03-advanced/     # 2 project advanced
    └── 04-expert/       # 2 project expert
```

## 🌱 Level 1: BEGINNER

### 📁 01-beginner/

- **01-hello-world/** - Program pertama dengan output dasar
- **02-variables/** - Deklarasi dan inisialisasi variabel
- **03-constants/** - Konstanta dan penggunaannya
- **04-data-types/** - Tipe data dasar (int, float, bool, string)
- **05-operators/** - Operator aritmatika, perbandingan, logika
- **06-control-flow/** - If-else, switch statement
- **07-loops/** - For loop dengan berbagai variasi
- **08-functions/** - Fungsi, parameter, return value
- **09-arrays-slices/** - Array dan slice operations
- **10-maps/** - Map untuk struktur key-value

### 🎯 Target Pembelajaran:

- Memahami syntax dasar Golang
- Menguasai tipe data dan variabel
- Menggunakan control flow dan perulangan
- Membuat dan menggunakan fungsi
- Bekerja dengan struktur data dasar

## 🚀 Level 2: INTERMEDIATE

### 📁 02-intermediate/

- **01-structs/** - Struct, methods, dan embedded struct
- **02-pointers/** - Pointer dan memory management
- **03-interfaces/** - Interface dan polymorphism
- **04-error-handling/** - Error handling patterns
- **05-packages/** - Package management dan modularitas
- **06-json-handling/** - JSON encoding/decoding
- **07-file-io/** - File operations dan I/O

### 🎯 Target Pembelajaran:

- Memahami OOP concepts di Go
- Menguasai interface dan polymorphism
- Implementasi error handling yang baik
- Bekerja dengan package dan modularitas
- Manipulasi JSON dan file I/O

## ⚡ Level 3: ADVANCED

### 📁 03-advanced/

- **01-goroutines/** - Concurrency dengan goroutines
- **02-channels/** - Channel untuk komunikasi antar goroutines
- **03-context/** - Context untuk lifecycle management

### 🎯 Target Pembelajaran:

- Memahami concurrency patterns
- Menguasai goroutines dan channels
- Implementasi context untuk cancelation
- Advanced synchronization techniques

## 🏆 Level 4: EXPERT

### 📁 04-expert/

- **01-reflection/** - Runtime type introspection
- **02-generics/** - Generic programming (Go 1.18+)
- **03-performance/** - Performance optimization dan profiling

### 🎯 Target Pembelajaran:

- Memahami reflection untuk metaprogramming
- Menguasai generics untuk code reusability
- Optimisasi performa dan memory management
- Advanced Go patterns dan best practices

## � Examples - Project Praktis

Selain learning examples, tersedia juga project lengkap untuk praktik langsung:

### 🟢 **Beginner Projects**

**1. Calculator CLI** (`examples/01-beginner/calculator-cli/`)

- **Struktur**: Single file dengan modular functions
- **Fitur**: Operasi matematika dasar (+, -, *, /, %, ^)
- **Tech**: Standard library, command-line interface
- **Goal**: Memahami functions, switch case, user input

**2. Todo List** (`examples/01-beginner/todo-list/`)

- **Struktur**: Struct-based dengan file storage
- **Fitur**: Add, list, complete, delete tasks
- **Tech**: JSON handling, file I/O operations
- **Goal**: Memahami structs, slices, file operations

### 🟡 **Intermediate Projects**

**3. REST API** (`examples/02-intermediate/rest-api/`)

- **Struktur**: MVC pattern dengan handlers
- **Fitur**: CRUD operations, JSON API
- **Tech**: Gorilla Mux, HTTP server, middleware
- **Goal**: Memahami HTTP, JSON, routing

**4. File Manager** (`examples/02-intermediate/file-manager/`)

- **Struktur**: Web interface dengan backend API
- **Fitur**: Upload, download, directory listing
- **Tech**: HTML templates, file operations, web server
- **Goal**: Memahami web development, templates

### 🟠 **Advanced Projects**

**5. Web Scraper** (`examples/03-advanced/web-scraper/`)

- **Struktur**: Package-based architecture
- **Fitur**: Concurrent scraping, data extraction
- **Tech**: GoQuery, concurrency, worker pools
- **Goal**: Memahami concurrency, HTML parsing

**6. Chat Server** (`examples/03-advanced/chat-server/`)

- **Struktur**: Microservice architecture
- **Fitur**: Real-time messaging, multiple rooms
- **Tech**: WebSocket, goroutines, hub pattern
- **Goal**: Memahami real-time communication, WebSocket

### 🔴 **Expert Projects**

**7. API Gateway** (`examples/04-expert/api-gateway/`)

- **Struktur**: Enterprise-level architecture
- **Fitur**: Load balancing, rate limiting, monitoring
- **Tech**: Redis, Prometheus, reverse proxy
- **Goal**: Memahami microservices, production patterns

**8. Performance Monitor** (`examples/04-expert/performance-monitor/`)

- **Struktur**: Real-time monitoring system
- **Fitur**: System metrics, WebSocket streaming, dashboard
- **Tech**: gopsutil, WebSocket, Prometheus, Charts
- **Goal**: Memahami system programming, real-time data

## 🏗️ Architecture Patterns

### **Project Structure Evolution**

```
Beginner:    main.go (single file)
             ↓
Intermediate: main.go + handlers/ (multiple files)
             ↓
Advanced:    cmd/ + internal/ + pkg/ (package-based)
             ↓
Expert:      full microservice architecture
```

### **Key Technologies per Level**

- **Beginner**: Standard library, basic I/O
- **Intermediate**: HTTP server, JSON, templates
- **Advanced**: Concurrency, WebSocket, external packages
- **Expert**: Monitoring, caching, production deployment

## 🚀 Cara Menggunakan

### **Learning Examples**

```bash
# Install Go (minimal versi 1.18)
go version
```

### **Project Examples**

```bash
# Masuk ke folder project yang diinginkan
cd examples/01-beginner/calculator-cli

# Install dependencies (jika ada)
go mod tidy

# Jalankan project
go run main.go
# atau untuk project dengan struktur cmd/
go run cmd/main.go
```

### **Development Flow**

```bash
# 1. Pelajari konsep di learning examples
cd 01-beginner/01-hello-world
go run main.go

# 2. Praktik dengan project examples  
cd examples/01-beginner/calculator-cli
go run main.go

# 3. Modifikasi dan eksperimen
# Edit kode sesuai kebutuhan

# 4. Lanjut ke level berikutnya
```

## 📚 Learning Path

### **Tahap 1: Foundation (Beginner)**

- ✅ Syntax dasar dan tipe data
- ✅ Control flow dan functions
- ✅ Arrays, slices, maps
- 🎯 **Project**: Calculator CLI, Todo List

### **Tahap 2: Intermediate Development**

- ✅ Structs, interfaces, pointers
- ✅ Error handling dan packages
- ✅ JSON dan file I/O
- 🎯 **Project**: REST API, File Manager

### **Tahap 3: Advanced Patterns**

- ✅ Goroutines dan channels
- ✅ Context dan synchronization
- ✅ Reflection dan testing
- 🎯 **Project**: Web Scraper, Chat Server

### **Tahap 4: Expert Level**

- ✅ Generics dan performance
- ✅ Production patterns
- ✅ Monitoring dan deployment
- 🎯 **Project**: API Gateway, Performance Monitor

## � Tips Pembelajaran

### **Beginner Level**

1. **Pelajari secara berurutan** - Mulai dari `01-hello-world`
2. **Praktik langsung** - Jalankan setiap contoh kode
3. **Modifikasi kode** - Coba ubah nilai dan lihat hasilnya
4. **Pahami error** - Baca pesan error dengan teliti
5. **Build simple projects** - Mulai dengan Calculator CLI

### **Intermediate Level**

1. **Fokus pada konsep** - Pahami struct, interface, error handling
2. **Buat project kecil** - Implementasikan REST API sederhana
3. **Baca dokumentasi** - Pelajari standard library packages
4. **Code review** - Analisis struktur project yang ada
5. **Web development** - Praktik dengan File Manager project

### **Advanced Level**

1. **Pelajari concurrency** - Pahami goroutines dan channels
2. **Praktik patterns** - Implementasikan design patterns
3. **Real-time apps** - Build Chat Server dengan WebSocket
4. **Performance** - Benchmark dan optimisasi kode
5. **External packages** - Integrasikan third-party libraries

### **Expert Level**

1. **Production ready** - Deploy aplikasi ke production
2. **Monitoring** - Implementasikan metrics dan logging
3. **Microservices** - Arsitektur terdistribusi
4. **Contribution** - Kontribusi ke open source projects
5. **Leadership** - Mentor dan design system architecture

## 🎯 Milestone Targets

### **Week 1-2: Go Fundamentals**

- ✅ Complete all beginner examples (01-10)
- ✅ Build Calculator CLI project
- ✅ Build Todo List project

### **Week 3-4: Intermediate Skills**

- ✅ Master structs, interfaces, error handling
- ✅ Build REST API project
- ✅ Build File Manager project

### **Week 5-6: Advanced Concepts**

- ✅ Understand concurrency patterns
- ✅ Build Web Scraper project
- ✅ Build Chat Server project

### **Week 7-8: Expert Level**

- ✅ Performance optimization techniques
- ✅ Build API Gateway project
- ✅ Build Performance Monitor project

## 🛠️ Development Tools

### **Essential Tools**

- **Go**: Version 1.21+ (latest stable)
- **Editor**: VS Code with Go extension
- **Version Control**: Git untuk source control

### **Development Tools**

- **Debugging**: Delve debugger (`dlv`)
- **Testing**: Built-in `go test` framework
- **Formatting**: `gofmt`, `goimports` for code formatting
- **Linting**: `golangci-lint` untuk code quality

### **Advanced Tools**

- **Profiling**: `go tool pprof` untuk performance analysis
- **Benchmarking**: `go test -bench` untuk performance testing
- **Documentation**: `godoc` untuk documentation generation

## 🌐 Production Deployment

### **Containerization**

```dockerfile
# Dockerfile example untuk Go apps
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

### **Docker Compose untuk Development**

```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "8080:8080"
  redis:
    image: redis:alpine
    ports:
      - "6379:6379"
```

## � Performance Benchmarks

Setiap project dilengkapi dengan benchmark tests:

```bash
# Run benchmarks
go test -bench=. -benchmem

# CPU profiling
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof

# Memory profiling  
go test -bench=. -memprofile=mem.prof
go tool pprof mem.prof
```

## 🤝 Contributing

Jika ingin berkontribusi ke learning repository ini:

1. Fork repository
2. Create feature branch
3. Add examples or improve documentation
4. Submit pull request
5. Follow Go best practices dan coding standards

## 📚 Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Blog](https://blog.golang.org/)
- [A Tour of Go](https://tour.golang.org/)

## 🤝 Kontribusi

Silakan berkontribusi untuk meningkatkan materi pembelajaran ini :D

## 🏁 Conclusion

Repository ini menyediakan learning path yang komprehensif untuk menguasai Golang dari level beginner hingga expert. Dengan kombinasi **23 learning examples** dan **8 production-ready projects**, Anda akan mendapatkan pengalaman pembelajaran yang holistik.

### **Key Takeaways**

✅ **Progressive Learning** - Dari basic syntax hingga production patterns
✅ **Hands-on Practice** - Real-world projects di setiap level
✅ **Modern Architecture** - Package structure yang scalable
✅ **Production Ready** - Deployment dan monitoring tools
✅ **Best Practices** - Go idioms dan coding standards

### **Next Steps**

1. **Start Learning** - Mulai dari `01-beginner/01-hello-world`
2. **Build Projects** - Praktik dengan examples projects
3. **Join Community** - Bergabung dengan Go developer community
4. **Keep Coding** - Practice makes perfect!

---

## 📄 License

Materi pembelajaran ini tersedia untuk tujuan edukasi dan open source development.

---

## 🎉 Happy Coding!

> _"The best way to learn Go is by writing Go code"_ - Rob Pike

**Keep coding, keep learning, keep growing! 🚀**
