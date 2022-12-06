excel-proc 处理 excel 数据，通过对数据的条件判断，使用模版内容组装自己想要的内容结果。
该工具的灵感来源于处理 “流水账单” 转化为 beancount 交易文件。

# 快速开始
## 例子
编写模版配置文件，文件 `config.tpl` 内容如下：
```
; ({col4} < 0) and ({col3} =~ "支付宝")
2022-{{substr 0 2 .col2}}-{{substr 3 5 .col2}} *  "{{.col3}}"
  Expenses:Live {{.col4}} CNY
  Liabilities:CreditCard:CMB {{mul .col4 -1}} CNY
```

* 第一行，必须 “;” 开头，定义 excel 数据的筛选条件。
    1. `{col4} < 0` 表示第四列的数据小与 0;
    2. `{col3} =~ "支付宝"` 表示第三列的数据包含 “支付宝” 字样;
    3. `and` 表示这两个条件都匹配成功;
    4. [更多条件](#条件)

* 其余的是模版内容，定义处理后数据被组装好的样子，使用该模版必须 “;” 后的条件匹配成功。
    1. `{{.col1}}` 填充第一列数据;
    2. `{{.col2}}`、`{{.col3}}` 等等都表示填充数据;
    3. `2022-{{substr 0 2 .col2}}-{{substr 3 5 .col2}}` 截取第二列的数据，将类似 "03/22" 处理为 "2022-03-22";
    4. `{{mul .col4 -1}}` 其中 mul 是个函数，表示第四列的数据乘以 -1;
    5. [更多函数](http://masterminds.github.io/sprig/)。

更多的例子在 example 目录下。

## 安装 excel-proc
```shell
# Go 1.15 或更低版本
go get -u github.com/miaogaolin/excel-proc

# Go 1.16 或更高版本
go install github.com/miaogaolin/excel-proc@latest
```

## 运行
```
excel-proc --config example/config.tpl example/data.csv
```
默认会生成 `default.bean` 文件，可以通过 `--output` 参数指定输出文件名。

在输出控制台也会出现如下信息：
```shell
[Warning] {col4} strconv.ParseFloat: parsing "RMB AMOUNT": invalid syntax: parse data error, condition=" ({col4} < 0) and ({col3} =~ "支付宝")", data=[SOLD POSTED DESCRIPTION RMB AMOUNT CARD NO(Last 4digits) Original Tran Amount]
```
警告：表示条件和数据之间是不能进行合理验证的，可忽略。

```shell
total count: 11
ignore count: 9
success count: 2
```
* total count:总共扫描 excel 的行数。
* ignore count: 和条件不匹配，或者出现如上的 Warning 信息的总数量。
* success count: 成果匹配的条件数。

# 数据处理
excel-proc 尽可能不修改 excel 数据内容，以增加不同数据的适配性，如下注意:

* 为了条件中支持数字比较，会将类似数字的字符串，统一处理为标准数字类型，例如： "-3,036.50"(有引号) 处理为 -3036.5, 去掉了引号和逗号。
* 不支持排序，所以需要提前处理好。

# 配置文件
在上面的例子中，编写了一条完整的规则，倘若需要多对条件和模版，则只需要通过空行分开，如下：  

```
; 条件1
模版1

; 条件2
模版2
```

1. 如果 “条件1” 匹配成功，则直接使用 “模版1” 内容渲染数据;
2. 如果 "条件1” 不成功则继续往下匹配查找，直到条件匹配成功；
3. 如果不想使用条件，则可以省略，代表所有数据都是用一个模版内容。

```
模版1
```
这个就只要一条模版内容，省略了条件，即时 “模版1” 后面空格相隔了多条 “模版” 内容也是不会应用，无效。

# 条件

* `==` 两值相等，例: `{col1} == "hello"`
* `!=` 两值不相等，例: `{col1} != "hello"`
* `>`、`<` 数字判断，例：`{col1} > 1`
* `in` 存在数组中，例: `{col1} in ["hello", "world"]`
* `not in` 不存在数组中，例: `{col1} not in ["hello", "world"]`
* `=~` 存在字符串中，例：`{col1} =~ "h"`
* `!~` 不存在字符串中，例：`{col1} !~ "h"`

如上的比较语句，之间需要验证如下：

* `and` 与，表示两边的条件都满足
* `or` 或，表示两边的条件满足其中一个即可

使用小括号，调整优先级：
```shell
# 表示验证 {col2} == 2 and {col3} == 3
{col1} == 1 or ({col2} == 2 and {col3} == 3)
```


# 模版
模版的语法就是原版的 Go 语言模版语法，例如想增加 if 语句，也是可以的，只不过模版在此基础上扩展了 [sprig](http://masterminds.github.io/sprig/) 相关函数。
如果在多个模板中有很多相同的内容，可以设置一个公共模板内容。

```
field1: "公共模板内容"

; 条件1
{{.field1}}
模板1

; 条件2
{{.field1}}
模板2
```
`field1` 就是公共的模板内容，这个字段名自定义，可使用 "双引号" 或 "单引号" 赋值。