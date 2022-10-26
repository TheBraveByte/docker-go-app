# Multi-step docker build

# Build phase
FROM golang:latest as builder

WORKDIR /usr/app

COPY go.mod go.sum ./

RUN go mod download 

COPY ./ /usr/app/

RUN go build -o main

# CMD [ "./main" ]

# Run phase -- copy over the result of the build phase to the 
# run phase below
FROM nginx

COPY --from=builder /usr/app /usr/share/nginx/html