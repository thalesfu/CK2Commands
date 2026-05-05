# CK2Commands

《十字军之王 II》存档辅助主程序。读取最新存档，写入图数据库（Neo4j / Nebula），并生成可在游戏控制台直接执行的脚本命令。

## 编译

```bash
go build -o ck2 ./main/main.go
```

## 用法

```bash
./ck2                    # 默认：读存档 → 写 Neo4j → 生成 auto.txt / cure.txt
./ck2 -w                 # 监视模式：存档变化时自动重新处理
./ck2 -f                 # 强制重载：不检查存档是否变化
./ck2 -target=nebula     # 写入 Nebula
./ck2 -target=both       # 同时写入两个数据库
./ck2 -bs                # 构建静态数据（文化/宗教/特性等）
./ck2 -m                 # 生成婚配建议脚本
./ck2 -wr                # 生成批量改宗（道教）脚本
```

完整参数说明和工作流程见 [docs/GUIDE.md](docs/GUIDE.md)。

## 目录结构

```
main/           入口
family/         各核心家族游戏逻辑
ck2/            存档路径 / SHA256 检测
people/         人物查询辅助
scriptbuilder.go / scriptgenerator.go   控制台命令生成
historypeople/  历史人物静态数据（自动生成）
historydynasty/ 历史王朝静态数据（自动生成）
earth/          头衔分配定义
docs/           用户文档
```
