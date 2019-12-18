# golu
some useful command-line tools

## All tools
- [dict](dict): Chinese & English translation
- [base64](base64)：base64 encode & decode
- [weather](): the current weather

### dict
Chinese-English translation tool
```
# onkey install
wget https://github.com/shiniao/golu/dict/dict.go && go build -o /usr/local/bin/dict dict.go

# how to use
➜ ~ dict hello
*******英汉翻译:*******
你好
*******网络释义:*******
Hello : 你好 / 您好 / 哈啰 / 喂 / 
Hello Kitty : 凯蒂猫 / 昵称 / 吉蒂猫 / 匿称 / 
Hello Bebe : 哈乐哈乐 / 乐扣乐扣 / 

~ dict 你好
```

### base64
base64 encode & decode tool
```shell script
# onkey install
wget https://github.com/shiniao/golu/base64/base64.go && go build -o /usr/local/bin/base64 base64.go

# how to use
➜ ~ base64 [-e] [-d] hello
# get the help
➜ ~ base64 -h

```

### weather
the current weather tool
```shell script
# onkey install
wget https://github.com/shiniao/golu/base64/base64.go && go build -o /usr/local/bin/base64 base64.go

# how to use
➜ ~ base64 [-e] [-d] hello
# get the help
➜ ~ base64 -h
```

