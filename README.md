This project's structure is:

```
<PROJECT_ROOT>
├── main.go
|
├── .gogradle
│   └── project_gopath
│       └── src
│           └── github.com
│               └── blindpirate
│                   └── gogland-issue -> ../../../../.. (a symbolic link pointing to <PROJECT_ROOT>)
|
└── vendor
    └── github.com
        └── urfave
            └── cli
                └── ...
   
```

Now I run:

```
export GOPATH=<PROJECT_ROOT>/.gogradle/project_gopath

go build github.com/blindpirate/gogland-issue
```

It works well since `GOPATH/src/github.com/blindpirate/gogland-issue` can be found.

But if I open `<PROJECT_ROOT>/.gogradle/project_gopath/src/github.com/blindpirate/gogland-issue` and set project GOPATH to '<PROJECT_ROOT>/.gogradle/project_gopath' in gogland, something doesn't work as follows:

![1](https://raw.githubusercontent.com/blindpirate/gogland-issue/master/issue.png)

My environment is 
```
Gogland (1.0 Preview) 1.0 EAP
Build #GO-171.4694.61, built on June 27, 2017
Gogland EAP User
Expiration date: September 25, 2017
JRE: 1.8.0_112-release-736-b21 x86_64
JVM: OpenJDK 64-Bit Server VM by JetBrains s.r.o
Mac OS X 10.12.1
```


