# Go 1.23 kullan
FROM golang:1.23

# Çalışma dizini oluştur
WORKDIR /app

# Go mod dosyalarını kopyala ve modülleri indir
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Geri kalan dosyaları kopyala
COPY . .

# Uygulamayı derle
RUN go build -o main .

# Uygulamayı başlat
CMD ["./main"]
