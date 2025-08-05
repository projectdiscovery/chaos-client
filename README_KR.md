<h1 align="center">
Chaos Client
</h1>
<h4 align="center">Chaos 데이터셋 API와 통신하기 위한 Go 클라이언트입니다.</h4>

<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/projectdiscovery/chaos-client">
<a href="https://github.com/projectdiscovery/chaos-client/releases"><img src="https://img.shields.io/github/downloads/projectdiscovery/chaos-client/total">
<a href="https://github.com/projectdiscovery/chaos-client/graphs/contributors"><img src="https://img.shields.io/github/contributors-anon/projectdiscovery/chaos-client">
<a href="https://github.com/projectdiscovery/chaos-client/releases/"><img src="https://img.shields.io/github/release/projectdiscovery/chaos-client">
<a href="https://discord.gg/projectdiscovery"><img src="https://img.shields.io/discord/695645237418131507.svg?logo=discord"></a>
<a href="https://twitter.com/pdchaos"><img src="https://img.shields.io/twitter/follow/pdchaos.svg?logo=twitter"></a>
</p>

<p align="center">
  <a href="https://github.com/projectdiscovery/chaos-client/blob/main/README.md">English</a> •
  <a href="https://github.com/projectdiscovery/chaos-client/blob/main/README_CN.md">中文</a> •
  <a href="https://github.com/projectdiscovery/chaos-client/blob/main/README_KR.md">한국어</a>
</p>

## 📦 설치 방법

```bash
go install -v github.com/projectdiscovery/chaos-client/cmd/chaos@latest
```

## 사용법

```bash
chaos -h
```

도움말이 출력되며, 다음과 같은 옵션들을 지원합니다:

```console
   -key string                  ProjectDiscovery Cloud (pdcp) API 키
   -d string                    서브도메인을 조회할 도메인
   -count                       지정한 도메인의 통계 출력
   -silent                      출력 최소화
   -o string                    출력 결과를 파일로 저장 (선택 사항)
   -dL string                   서브도메인을 조회할 도메인 목록 파일 (선택 사항)
   -json                        JSON 형식으로 출력
   -version                     Chaos 버전 출력
   -v, -verbose                 상세 출력 모드
   -up, -update                 Chaos를 최신 버전으로 업데이트
   -duc, -disable-update-check 자동 업데이트 확인 비활성화
```

또한, bash 프로파일에 API 키를 환경 변수로 설정할 수 있습니다:

```bash
export PDCP_API_KEY=xxxxx
```

## `PDCP_API_KEY` 발급 방법

[cloud.projectdiscovery.io](https://cloud.projectdiscovery.io?ref=api_key)에서 회원가입 또는 로그인 후 API 키를 발급받을 수 있습니다.

## Chaos 실행 예시

도메인의 서브도메인을 조회하려면 아래 명령어를 실행하세요:

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

💡 참고사항
-----

- **API는 IP 기준 분당 60회로 속도 제한됩니다.**
- Chaos API는 **도메인 이름 기반 조회만 지원**합니다.

## 라이브러리로 사용하기
`Chaos`는 CLI뿐 아니라 라이브러리 형태로도 사용 가능합니다. `Options` 구조체를 통해 설정 후 호출합니다.

### 예제 코드
```go
package main

import (
	"os"
	"github.com/projectdiscovery/chaos-client/internal/runner"
	"github.com/projectdiscovery/chaos-client/pkg/chaos"
)

func main() {
	var results []chaos.Result
	opts := &runner.Options{
		Domain: "projectdiscovery.io",
		APIKey: os.Getenv("PDCP_API_KEY"),
		OnResult: func(result interface{}) {
			if val, ok := result.(chaos.Result); !ok {
				results = append(results, val)
			}
		},
	}

	runner.RunEnumeration(opts)
}

```
💡 주의

위 프로그램을 실행하려면 `PDCP_API_KEY` 환경 변수가 설정되어 있어야 합니다.

👨‍💻 커뮤니티
-----

[Discord 커뮤니티](https://discord.gg/projectdiscovery)에 참여하거나  
[Twitter (@pdchaos)](https://twitter.com/pdchaos)를 팔로우하여 최신 정보를 받아보세요.

여러분의 기여 덕분에 커뮤니티는 계속 성장하고 있습니다. 감사합니다! ❤️
