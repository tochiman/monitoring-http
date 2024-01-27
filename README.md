# monitoring-http

## Development-Environment
|      名前      |           バージョン          |     説明     |
|:--------------:|:-----------------------------:|:------------:|
|     Ubuntu     | 22.04.3 LTS (Jammy Jellyfish) |      OS      |
|       go       |     go1.18.1 (linux/amd64)    | 開発中に使用 |
|  Docker-Engine |             24.0.7            |  Docker関連  |
|   containerd   |             1.6.25            |  Docker関連  |
|      runc      |             1.1.10            |  Docker関連  |
|   docker-init  |             0.19.0            |  Docker関連  |
| Docker Compose |            v2.21.0            |       -      |

## installation
1. Download
```bash
git pull https://github.com/tochiman/monitoring-http.git
```
2. Configure monitoring-http/app/targets.yml
```yml
targets:
  Google:
    Namespace: "test"
    Name: "Google"
    Help:  "Status of Google Service"
    URL: "https://google.com/"
  Youtube:
    Namespace: "test"
    Name: "Youtube"
    Help:  "Status of Youtube Service"
    URL: "https://youtube.com/"
```

3. 