# gotbot

Project represents a simple discord bot written in go language

## build:

```bash
go build
```

## configuration:

You can use different configuration sources.
```bash
./gotbot -ct {cmd, yaml, env}
```

## run:

```bash
./gotbot -t BOT_TOKEN
```

## commands:

* ### !help return reference
    ```text
    @gotbot !help
    ```
* ### !rnd 
    rnd returns random number
    with max value border
    ```text
    @gotbot !rnd 10
    ```
    with min and max value borders
    ```text
    @gotbot !rnd 50 100
    ``` 
* ### !roll 
    roll returns random element from provided list
    ```text
    @gotbot !roll red green blue 
    ```

## heroku

for debugging purposes you can create a .env file (key=value):
```text
TOKEN=your_token
```
and run heroku with
```bash
heroku local -e ${ENV.FILE.PATH}
```