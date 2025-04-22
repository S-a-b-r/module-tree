## Утилита для отрисовки дерева зависимостей с определенной глубиной в консоли

```shell
    go build -o=tree-gen cmd/go-mod-file-analyze/main.go
```

```shell
    ./tree-gen --prefix=bg --depth=2 --substring=gitlab.globars.ru
```