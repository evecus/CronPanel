# CronPanel — Crontab 管理面板

现代化的 Linux Crontab 定时任务 Web 管理工具。

## 功能特性

- 📋 **查看任务** — 实时查看所有 crontab 任务
- ➕ **添加任务** — 多种时间模式（每天/每N天/每周/每月/自定义）
- 🗑️ **删除任务** — 一键删除定时任务
- ⏸️ **启用/停用** — 无需删除即可禁用任务
- 📝 **脚本支持** — 直接命令、脚本路径、或在线编写 sh 脚本内容
- 🌐 **现代化 UI** — 深色主题，响应式设计

## 编译

需要 Go 1.18 或更高版本。

### 快速编译

\`\`\`bash
chmod +x build.sh
./build.sh
\`\`\`

### 手动编译

\`\`\`bash
# Linux amd64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o cronpanel-linux-amd64 .

# Linux arm64
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o cronpanel-linux-arm64 .
\`\`\`

### Docker 编译（无需本地 Go 环境）

\`\`\`bash
# 使用 Docker 编译两个架构的二进制文件
docker run --rm -v "$PWD":/app -w /app golang:1.21-alpine sh -c "
  GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o dist/cronpanel-linux-amd64 . &&
  GOOS=linux GOARCH=arm64 go build -ldflags='-s -w' -o dist/cronpanel-linux-arm64 .
"
\`\`\`

## 运行

\`\`\`bash
# 赋予执行权限
chmod +x cronpanel-linux-amd64

# 默认端口 8899
./cronpanel-linux-amd64

# 自定义端口
PORT=9090 ./cronpanel-linux-amd64
\`\`\`

打开浏览器访问: http://localhost:8899

## 注意事项

- 程序需要有权限执行 \`crontab\` 命令
- 编写的 Shell 脚本默认保存到 \`/tmp/crontab-manager-scripts/\` 目录
- 建议以执行 crontab 任务的用户身份运行

## 文件说明

\`\`\`
crontab-manager/
├── main.go       # 主程序（HTTP 服务器 + Crontab 操作）
├── html.go       # 前端 HTML/CSS/JS（嵌入在 Go 中）
├── go.mod        # Go 模块定义
├── build.sh      # 一键编译脚本
└── README.md     # 说明文档
\`\`\`
