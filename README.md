IsUp provides a simple API that allows you to check if a service is running or not.

Compile
-------
Check out the sources and run

    go get -d ./src/main/ && go build -o ./bin/IsUp ./src/main
    
Run
---

    ./bin/IsUp
    
By default IsUp will listen on port 8888. If you want to change it, pass the "--port" option to the program.

    ./bin/IsUp --port=3333
    
IsUp will wait up to 10 second before declaring the host down. 
If you want to change this duration, pass the "--default-timeout" option:

    ./bin/IsUp --default-timeout=30
    
Use
---

You can check the status of any host and port by accessing the URL:
    
    http://localhost:8888/{host}/{port}
    
For example, to check the GitHub site go to:

    http://localhost:8888/github.com/80
    
If a TCP connection can be established, the service will return:

    {
        "success": true
    }
    
Otherwise the response will be:

    {
        "success": false
    }
    
If you want to use a timeout value other than the default, pass a "timeout"
parameter to the request:

    http://localhost:8888/github.com/80?timeout=5
    