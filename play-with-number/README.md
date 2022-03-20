# play-with-number
https://womenwhogotokyo.github.io/codelab/play-with-number/

## Build binary
```bash
$ go build -v -o play-with-number
```

## Run the game with `todayis` mode
```bash
$ ./play-with-number -game todayis
20210418は素数ではありません
```

## Run the game with `prime` mode
```bash
$ ./play-with-number -game prime
数字を入力してください
5
5は素数です
```

## Run the test of the prime mode
```bash
$ go test ./prime
```

[![Open in Cloud Shell](http://gstatic.com/cloudssh/images/open-btn.svg)](https://console.cloud.google.com/cloudshell/editor?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2FWomenWhoGoTokyo%2Fcodelab&cloudshell_workspace=play-with-number)
