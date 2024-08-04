# Performance comparison of assignment and append in slice

Running example:

```bash
GORESEARCH_SIZE=1000 go test -bench=.* -benchmem ./...
```

Result example:

```bash

```

Running example without optimizations:

```bash
GORESEARCH_SIZE=1000 go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash

```
