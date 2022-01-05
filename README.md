Decode AAC (broadcasted on RTMP) to PCM on the fly with FDKAAC

## Run server
```bash
$ go run main.go
```

## Broadcast to server
```bash
$ ffmpeg -re -i file.mp4 -vn -c:a copy -f flv rtmp://localhost
```

