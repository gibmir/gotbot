# gotbot

Project represents a simple discord bot written in go language

## build:

```bash
go build
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
* ### !rnd return random number
    with max value border
    ```text
    @gotbot !rnd 10
    ```
    with min and max value borders
    ```text
    @gotbot !rnd 50 100
    ``` 
* ### !roll return random element from provided list
    ```text
    @gotbot !roll red green blue 
    ```