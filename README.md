## Abstract
This is a trello CLI library.
You can confirm cards by CLI.

## Use Case
- Confirm Cards(Tasks)
- Posting Card(Task)

## How To Use

#### build single binary
download binary directory (comming soon...)
```
$ make build
```

#### prepare configuration file (config.toml)
```
make setup
```

#### fill in configuration file
edit configration file.
generate Developer API Keys from following URL.
https://trello.com/app-key
```
$ vim config.toml
```

#### run command
```
$ ./trellocli  -config ./config.toml list
# display cards

$ ./trellocli -config ./config.toml add -card_name "Buy Milk"
# add card
```
