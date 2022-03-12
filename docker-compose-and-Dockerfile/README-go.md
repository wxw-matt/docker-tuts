# Note
The code is originally from https://github.com/dexterdarwich/ws-compare/tree/master/go

## Endpoints

| Route                    | Request example  | Response example                  |
|:-------------------------|:-----------------|:----------------------------------|
| GET /hello               | /hello           | {"message":"Hello, World"}        |
| GET /greeting/{name}     | /greeting/Jane   | {"greeting":"Hello, Jane!"}       |
| GET /fibonacci/{number}  | /fibonacci/45    | {"input":45,"output":1134903170}  |

### Build
    go build -o webservice

  The compiled program should now be in the `ws-compare/go` folder:

    webservice

### Run
    ./webservice

# Notes:
- The code is not suitable for production use. And is not idiomatic to each language. Also, lacks tests and documentation.
- This project was written to compare some metrics using different programming languages. The metrics I was looking for:
    - Compiled artifact size
    - Size in memory while running idle
    - Size in memory while running at peak requests per second
    - Requests per second, latency, and timeouts
    - Packaged, deployment ready build with the relevant runtime. (Docker image)
    - Compile/build time
