<h1 align="center">
Chaos Client
</h1>
<h4 align="center">与 Chaos 数据集 API 通信的 Go 客户端。</h4>

<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/projectdiscovery/chaos-client">
<a href="https://github.com/projectdiscovery/chaos-client/releases"><img src="https://img.shields.io/github/downloads/projectdiscovery/chaos-client/total">
<a href="https://github.com/projectdiscovery/chaos-client/graphs/contributors"><img src="https://img.shields.io/github/contributors-anon/projectdiscovery/chaos-client">
<a href="https://github.com/projectdiscovery/chaos-client/releases/"><img src="https://img.shields.io/github/release/projectdiscovery/chaos-client">
<a href="https://discord.gg/projectdiscovery"><img src="https://img.shields.io/discord/695645237418131507.svg?logo=discord"></a>
<a href="https://twitter.com/pdchaos"><img src="https://img.shields.io/twitter/follow/pdchaos.svg?logo=twitter"></a>
</p>


## 安装

```bash
go install -v github.com/projectdiscovery/chaos-client/cmd/chaos@latest
```

## 使用

```bash
chaos -h
```

这将会显示工具的帮助说明。以下是它支持的所有选项。

| 选项                       | 描述                              | 示例                                                    |
|----------------------------|------------------------------------------|------------------------------------------------------------|
| `-d`                       | 要查找子域的域            | `chaos -d uber.com`                                        |
| `-count`                   | 显示指定域的统计信息 | `chaos -d uber.com -count`                                 |
| `-o`                       | 将输出写入文件（可选）       | `chaos -d uber.com -o uber.txt`                            |
| `-json`                    | 将输出打印为 json 格式                     | `chaos -d uber.com -json`                                  |
| `-key`                     | Chaos API 的密钥                        | `chaos -key API_KEY`                                       |
| `-dL`                      | 带有域列表的文件（可选）     | `chaos -dL domains.txt`                                    |
| `-silent`                  | 让输出更加简洁                   | `chaos -d uber.com -silent`                                |
| `-version`                 | 打印 Chaos 客户端的当前版本    | `chaos -version`                                           |
| `-verbose`                 | 显示详细输出                      | `chaos -verbose`                                           |
| `-update`                  | 更新到最新版本                | `chaos -up`                                                | 
| `-disable-update-check`    | 禁用自动更新检查          | `chaos -duc`                                               |

您还可以在 bash profile 中将 API 密钥设置为环境变量。

```bash
export CHAOS_KEY=CHAOS_API_KEY
```

### 如何利用 `API_KEY`

Chaos DNS API 处于测试版，仅对被邀请的人可用。您可以在 [chaos.projectdiscovery.io](https://chaos.projectdiscovery.io) 请求邀请。

## 运行 Chaos

为了获取某个域的的子域，请使用以下命令。

```bash
chaos -d uber.com -silent

restaurants.uber.com
testcdn.uber.com
approvalservice.uber.com
zoom-logs.uber.com
eastwood.uber.com
meh.uber.com
webview.uber.com
kiosk-api.uber.com
utmbeta-staging.uber.com
getmatched-staging.uber.com
logs.uber.com
dca1.cfe.uber.com
cn-staging.uber.com
frontends-primary.uber.com
eng.uber.com
guest.uber.com
kiosk-home-staging.uber.com
```

💡 注意
-----

- **这个 API 的速率限制为 60 请求 / 分钟 / ip**
- Chaos API **只** 支持域名查询。

👨‍💻 社区
-----

欢迎您加入我们的 [Discord 社区](https://discord.gg/projectdiscovery)。您也可以在 [Twitter](https://twitter.com/pdchaos) 上关注我们，了解与 Chaos 项目相关的一切。


再次感谢您的贡献，让社区充满活力。 :heart:
