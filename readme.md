# ![GoLang](http://xiaorui.cc/static/golang_jump.gif)GoLang 
Go 亦称为 Golang (译注：按照 Rob Pike 说法，语言叫做 Go，Golang 只是官方网站的网址)，是由谷歌开发的一个开源的编译型的静态语言。


Golang 的主要关注点是使得高可用性和可扩展性的 Web 应用的开发变得简便容易。
## GoLang 环境部署
+ ### 安装 
    Golang 支持三个平台：Mac，Windows 和 Linux（译注：不只是这三个，也支持其他主流平台）,可以在[下载地址](https://golang.org/dl/)中下载相应平台的二进制文件。

    #### Mac OS
    在 [下载地址](https://golang.org/dl/) 下载安装程序。双击开始安装并且遵循安装提示，会将 Golang 安装到 /usr/local/go 目录下，同时 /usr/local/go/bin 文件夹也会被添加到 PATH 环境变量中。

    #### Windows
    在 [下载地址](https://golang.org/dl/) 下载 MSI 安装程序。双击开始安装并且遵循安装提示，会将 Golang 安装到 C:\Go 目录下，同时 c:\Go\bin 目录也会被添加到你的 PATH 环境变量中。
    * Win平台环境变量配置，可更改默认
    * __GOROOT :__ GoLang安装路径
    * __GOPATH :__ Go项目开发路径，即 **工作空间**
    * __Path :__ 除添加 go 安装根目录下 bin 路径外，可额外添加 **GOPATH** 下的 bin 路径，便于实时使用工作空间下的执行程序

    #### Linux
    + 在 [下载地址](https://golang.org/dl/) 下载 tar 文件，并解压到 /usr/local。
    + 添加 /usr/local/go/bin 到 PATH 环境变量中。Go 就已经成功安装在 Linux 上了。
+ ### Go环境配置
    使用 **go env** 命令查看Go环境配置，默认配置如下：
    ````js
        go env -w GO111MODULE=on                      //开启mod
        go env -w GOPROXY=https://goproxy.io,direct   //设置代理
    ````
    若出现 <font color='red'>go.mod file not found in current directory or any parent directory</font> 的编译问题，可执行一下配置
    >       go env -w GO111MODULE=auto
    若出现 <font color='green'>控制台中报错 package xxx is not in GOROOT</font> 的问题，可使用以下配置:
    >       go env -w GO111MODULE=off
    若出现 <font color='blue'>编译器报错 no required module provides package XXXXX</font>的问题，可使用如下操作：
    >       进入XXXXX第三方包根路径下，执行 go mod init 命令，初始化包后，重试。
    **eg:** 在 gopath 查找包，按照 __GOROOT__ 和 __GPPATH__ 目录下 src/xxx 依次查找。在 gomod 下查找包，解析 go.mod 文件查找包，mod 包名就是包的前缀，里面的目录就后续路径了。在 gomod 模式下，查找包就不会去 gopath 查找，只是 gomod 包缓存在 **{GOPATH}**/pkg/mod 里面。

    #### 如何使用 Modules 
    __GO111MODULE__ 有三个值：off, on（默认值）和auto。
    + __off__，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
    + __on__，go命令行会使用modules，不会去GOPATH目录下查找。
    + __auto__，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
        + 当前目录在GOPATH/src之外且该目录包含go.mod文件 
        + 当前文件在包含go.mod文件的目录下面
    
    __eg:__ auto,看当前目录或者上级目录是否存在go.mod,go.sum，如果存在，则启用go mod
    #### mod 命令
    |    命令    |   说明    |
    |   :-----:   |  :-----:   |
    |download|download modules to local cache(下载依赖包)|
    |edit|edit go.mod from tools or scripts(编辑go.mod)|
    |graph|print module requirement graph (打印模块依赖图)|
    |init|initialize new module in current directory(在当前目录初始化mod,生成go.mod文件)|
    |tidy|add missing and remove unused modules(拉取缺少的模块，移除不用的模块)|
    |vendor|make vendored copy of dependencies(将依赖复制到vendor下)|
    |verify|verify dependencies have expected content (验证依赖是否正确)|
    |why|explain why packages or modules are needed(解释为什么需要依赖)|
    
    __go.mod文件一旦创建后，它的内容将会被go toolchain全面掌控。go toolchain会在各类命令执行时，比如go get、go build、go mod等修改和维护go.mod文件__
+ ### 获取插件及工具包
    目前GoLang开发依赖第三方工具包及插件库，插件获取需要使用git远程仓库，国内请求git需要修改host文件。
    #### Win平台git配置步骤如下
    + 切换至C:\Windows\System32\drivers\etc，找到**hosts**文件，并开放写入权限
    + 在**hosts**文件内添加以下配置

            192.30.255.112 github.com git
            185.31.16.184 github.global.ssl.fastly.net
        **eg：** 使用ping命令测试 github.com ,若返回参数，则成功。

    #### 获取第三方包
    切换至**GOPATH**目录,创建bin、pkg、src路径。
    + src路径
        +  创建/golang.org/x/目录并在该目录下执行以下命令：

                git clone https://github.com/golang/tools.git
                git clone https://github.com/golang/lint.git
        + 安装插件
  
                go install github.com/ramya-rao-a/go-outline
                go install github.com/acroca/go-symbols
                go install golang.org/x/tools/cmd/guru
                go install golang.org/x/tools/cmd/gorename
                go install github.com/josharian/impl
                go install github.com/rogpeppe/godef
                go install github.com/sqs/goreturns
                go install github.com/golang/lint/golint
                go install github.com/cweill/gotests/gotests
                go install github.com/ramya-rao-a/go-outline
                go install github.com/acroca/go-symbols
                go install golang.org/x/tools/cmd/guru
                go install golang.org/x/tools/cmd/gorename
                go install github.com/josharian/impl
                go install github.com/rogpeppe/godef
                go install github.com/sqs/goreturns
                go install github.com/golang/lint/golint
                go install github.com/cweill/gotests/gotests
            <font color='red'>或</font>

                切换至src/golang.org/x/tools/cmd 路径，执行 go install ...
    + VSCode的扩展中安装go的插件
        **插件添加配置参数**
        ````json
        "[go]": { 
            "editor.formatOnSave": true, 
            "editor.codeActionsOnSave": { 
                "source.organizeImports": true, 
            },    // Optional: Disable snippets, as they conflict with completion ranking. 
            "editor.snippetSuggestions": "none", 
        },
        "[go.mod]": { 
            "editor.formatOnSave": true, 
            "editor.codeActionsOnSave": { 
                "source.organizeImports": true, 
            },
        },
        "gopls": {
            "usePlaceholders": true, 
            //"completeUnimported": true,
            "experimentalWorkspaceModule": true,
        },
        ````

        在git命令行工具执行命令（一定要在GOPATH指定的目录下执行）
        go get  // 将依赖项添加到当前模块并安装它们
        例如,若启用VSCode代码自动提示功能，需要安装gocode

            go get -u  github.com/mdempsky/gocode
            go get -u  github.com/go-delve/delve

        若尝试调试Go代码，可使用 **[配置教程](https://www.cnblogs.com/ljhoracle/p/11047083.html)** 的方式进行配置

## Go语法学习总结
+ ### if
    **if** 语句的语法是
        
        if condition {
        }

    如果 condition 为真，则执行 { 和 } 之间的代码
    **if** 还有另外一种形式，它包含一个 statement 可选语句部分，该组件在条件判断之前运行。它的语法是

        if statement; condition {
        }
    statement 一般放置条件初始化语句
    **eg:** else 语句应该在 if 语句的大括号 } 之后的同一行中。如果不是，编译器会不通过。__原因__ 是 Go 语言的分号是自动插入。你可以在这里阅读分号插入规则 https://golang.org/ref/spec#Semicolons
    __在 Go 语言规则中，它指定在 } 之后插入一个分号，如果这是该行的最终标记。因此，在 if 语句后面的 } 会自动插入一个分号。__
+ ### for
  **for** 是 Go 语言唯一的循环语句。Go 语言中并没有其他语言比如 C 语言中的 while 和 do while 循环。

        for initialisation; condition; post {
        }
    **initialisation** 只执行一次。循环初始化后，将检查循环**condition**。如果计算结果为 true ，则 {} 内的循环体将执行，接着执行 **post** 语句。**post** 语句将在每次成功循环迭代后执行。在执行 post 语句后，**condition** 将被再次检查。如果为 true, 则循环将继续执行, 否则 **for** 循环将终止。（译注：这是典型的 for 循环三个表达式，第一个为初始化表达式或赋值语句；第二个为循环条件判定表达式；第三个为循环变量修正表达式，即此处的 post ）

        package main

        import (
            "fmt"
        )

        func main() {
            for i := 1; i <= 10; i++ {
                fmt.Printf(" %d",i)
            }
        }
    **for** 循环的三部分，**initialisation**、**condition**、**post** 语句都是可选的。在以下的程序中，初始化语句和 post 语句都被省略了。i 在 for 循环外被初始化成了 0。只要 i<=10 循环就会被执行。在循环中，i 以 2 的增量自增。

        i := 0
        for ;i <= 10; { // initialisation and post are omitted
            fmt.Printf("%d ", i)
            i += 2
        }

    无限循环的语法是：

        for {
        }
    **range** 遍历
    类似 __Java__ 语法中的 __for...each__ 用法，range 返回索引和该索引处的值。
    ````go
        a := [...]float64{67.7, 89.8, 21, 78}
        sum := float64(0)
        for i, v := range a {//range returns both the index and value
            fmt.Printf("%d the element of a is %.2f\n", i, v)
            sum += v
        }
    ````
    如果你只需要值并希望忽略索引，则可以通过用 _ 空白标识符替换索引来执行。
+ ### switch
  switch 是一个条件语句，用于将表达式的值与可能匹配的选项列表进行比较，并根据匹配情况执行相应的代码块。它可以被认为是替代多个 if else 子句的常用方式。
  __在 Go 中，每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行__
    + **语法实例1：**
        ````go
            finger := 4
            switch finger {
            case 1:
                fmt.Println("Thumb")
            case 2:
                fmt.Println("Index")
            case 3:
                fmt.Println("Middle")
            case 4:
                fmt.Println("Ring")
            case 5:
                fmt.Println("Pinky")
            default: // 默认情况 不一定只能出现在 switch 语句的最后，它可以放在 switch 语句的任何地方。
                fmt.Println("incorrect finger number")
            }
        ````
    switch finger 将 finger 的值与每个 case 语句进行比较。通过从上到下对每一个值进行对比，并执行与选项值匹配的第一个逻辑。在上述样例中， finger 值为 4，因此打印的结果是 Ring 。**选项列表中，case 不允许出现重复项。**

    + **语法实例2：**
        ````go
            switch finger := 8; finger {
            case 1:
                fmt.Println("Thumb")
            case 2:
                fmt.Println("Index")
            case 3:
                fmt.Println("Middle")
            case 4:
                fmt.Println("Ring")
            case 5:
                fmt.Println("Pinky")
            default: // 默认情况
                fmt.Println("incorrect finger number")
            }
        ````
        switch finger：= 8; finger 中， 先声明了finger 变量，随即在表达式中使用了它。在这里，finger 变量的作用域仅限于这个 switch 内。
    + **无表达式的 switch**
        在 switch 语句中，表达式是可选的，可以被省略。如果省略表达式，则表示这个 switch 语句等同于 switch true，并且每个 case 表达式都被认定为有效，相应的代码块也会被执行。
        ````go
            num := 75
            switch { // 表达式被省略了
            case num >= 0 && num <= 50:
                fmt.Println("num is greater than 0 and less than 50")
            case num >= 51 && num <= 100:
                fmt.Println("num is greater than 51 and less than 100")
            case num >= 101:
                fmt.Println("num is greater than 100")
            }
        ````
    + **fallthrough 语句**
        使用 fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中。
        ````go
            switch num := 15 * 5; { // num is not a constant
            case num < 50:
                fmt.Printf("%d is lesser than 50\n", num)
                fallthrough
            case num < 100:
                fmt.Printf("%d is lesser than 100\n", num)
                fallthrough
            case num < 200:
                fmt.Printf("%d is lesser than 200", num)
            }
        ````
+ ### 数组
    Go 语言中不允许混合不同类型的元素，例如包含字符串和整数的数组。（如果是 interface{} 类型数组，可以包含任意类型）
    + 若数组仅声明，数组中的所有元素都被自动赋值为数组类型的零值。 
    + 数组初始化，不必填写数组大小，并用 ... 代替，让编译器为你自动计算长度，例如：
        ````go
            arr3 := [...]int{17, 18, 19, 20}
        ````
    + 多维数组声明格式 **[n][m]T**。
+ ### 切片
    切片是由数组建立的一种方便、灵活且功能强大的包装（Wrapper）。切片本身不拥有任何数据。它们只是对现有数组的**引用**。声明格式 **[ ]T**。
    + 案例1
        ````go
            a := [5]int{76, 77, 78, 79, 80}
            var b []int = a[1:4] // creates a slice from a[1] to a[3]
        ````
        语法 a[start:end] 创建一个从 a 数组索引 start 开始到 end - 1 结束的切片。因此，在上述程序的第 9 行中, a[1:4] 从索引 1 到 3 创建了 a 数组的一个切片表示。因此, 切片 b 的值为 [77 78 79]。
    + 案例2
        ````go
            c := []int{6, 7, 8} // creates and array and returns a slice reference
            fmt.Println(c)
        ````
        声明式创建切片
    
    **切片修改**
    对切片的更改反映在数组中。当多个切片共用相同的底层数组时，每个切片所做的更改将反映在数组中。
    ````go
        numa := [3]int{78, 79 ,80}
        nums1 := numa[:] // creates a slice which contains all elements of the array
        nums2 := numa[:]
        fmt.Println("array before change 1", numa)
        nums1[0] = 100
        fmt.Println("array after modification to slice nums1", numa)
        nums2[1] = 101
        fmt.Println("array after modification to slice nums2", numa)     
    ````
    在 9 行中，numa [:] 缺少开始和结束值。开始和结束的默认值分别为 0 和 len (numa)。
    输出是
    ````
        array before change 1 [78 79 80]
        array after modification to slice nums1 [100 79 80]
        array after modification to slice nums2 [100 101 80]
    ````
    **切片的长度和容量**
    切片的长度是切片中的元素数。切片的容量是从创建切片索引开始的底层数组中元素数。
    ````go
        fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
        fruitslice := fruitarray[1:3]
        fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
    ````
    
    在上面的程序中，fruitslice 是从 fruitarray 的索引 1 和 2 创建的。因此，fruitlice 的长度为 2。

    fruitarray 的长度是 7。fruiteslice 是从 fruitarray 的索引 1 创建的。因此, fruitslice 的容量是从 fruitarray 索引为 1 开始，也就是说从 orange 开始，该值是 6。因此, fruitslice 的容量为 6。
    <font color='red' size='4'>切片可以重置其容量。任何超出这一点将导致程序运行时抛出错误。</font>
    **函数创建：** make（[]T，len，cap）[]T 通过传递类型，长度和容量来创建切片。容量是可选参数, 默认值为切片长度。make 函数创建一个数组，并返回引用该数组的切片。
    **追加切片元素**
    数组的长度是固定的，它的长度不能增加。但切片是动态的，使用 append 可以将新元素追加到切片上。append 函数的定义是 func append（s[]T，x ... T）[]T。
    + __若切片初始容量已满，则追加后，容量翻番。__
    + __切片类型的零值为 nil。一个 nil 切片的长度和容量为 0。可以使用 append 函数将值追加到 nil 切片。切片容量为追加数据后初始值。__
    + 可以使用 ... 运算符将一个切片添加到另一个切片。例如：
        ````go
            veggies := []string{"potatoes", "tomatoes", "brinjal"}
            fruits := []string{"oranges", "apples"}
            food := append(veggies, fruits...)
        ````
+ ### 可变参数函数
    可变参数函数是一种参数个数可变的函数。
     __只有函数的最后一个参数才允许是可变的,函数最后一个参数被记作 ...T。__
    <font color='red' size='6' >...</font>俗称 __语法糖__。若用于数组声明，表示不定长数组；若用于函数声明，可表示 **不定参声明**  与 **切片传值**。
+ ### Map
    map 是在 Go 中将值（value）与键（key）关联的内置类型。通过相应的键可以获取到值。声明语法：map[type of key]type of value
    **函数创建：** make(map[type of key]type of value)
    ````go
        personSalary := make(map[string]int)
    ````
    <font color='red' size='4'>map 的零值是 nil。如果你想添加元素到 nil map 中，会触发运行时 panic。因此 map 必须使用 make 函数初始化。</font>
    <font color='green' size='4'>注：</font>__go与java语法存在类似对应关系，切片-->Array [Java],Map-->Map[ Java],nil-->null，都可使用make函数进行初始化。__

