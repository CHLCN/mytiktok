# mytiktok
类抖音短视频平台

## 运行前端
```bash
cd frontend

npm install

npm run dev
```

## 运行后端
```bash
cd backend

go build .

./backend.exe
```

## 给后端配置热更新
```bash
go install github.com/beego/bee/v2@latest

# 将$GOPATH/bin中的bee.exe放入GO的执行目录$GOROOT/bin

bee run
# 此时就不需每次改动后端代码后重新build
```
