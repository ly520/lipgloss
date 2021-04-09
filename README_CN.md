Lip Gloss
=======

<p>
    <img src="https://stuff.charm.sh/lipgloss/lipgloss-header-github.png" width="340" alt="Lip Gloss Title Treatment"><br>
    <a href="https://github.com/charmbracelet/lipgloss/releases"><img src="https://img.shields.io/github/release/charmbracelet/lipgloss.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/charmbracelet/lipgloss?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/charmbracelet/lipgloss/actions"><img src="https://github.com/charmbracelet/lipgloss/workflows/build/badge.svg" alt="Build Status"></a>
</p>

漂亮的终端布局的样式定义。建造时考虑了TUIs。

![Lip Gloss example](https://stuff.charm.sh/lipgloss/lipgloss-example.png)

Lip Gloss 采用了一种表达性的、陈述性的方法来进行终端渲染。

熟悉CSS的用户使用Lip Gloss会感到宾至如归。 

```go

import "github.com/charmbracelet/lipgloss"

var style = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
    PaddingTop(2).
    PaddingLeft(4).
    Width(22)

    fmt.Println(style.Render("Hello, kitty."))
```


## Colors

Lip Gloss 支持以下颜色配置：

### ANSI 16 colors (4-bit)

```go
lipgloss.Color("5")  // magenta 洋红
lipgloss.Color("9")  // red 红
lipgloss.Color("12") // light blue 浅蓝
```

### ANSI 256 Colors (8-bit)

```go
lipgloss.Color("86")  // aqua 湖绿色
lipgloss.Color("201") // hot pink 艳粉色
lipgloss.Color("202") // orange 橙色
```

### True Color (24-bit)

```go
lipgloss.Color("#0000FF") // good ol' 100% blue
lipgloss.Color("#04B575") // a green
lipgloss.Color("#3C3C3C") // a dark gray
```

终端的颜色配置文件将被自动检测，并且外部的颜色当前调色板的色域将自动强制到最接近的色域

可用值。


### 自适应颜色

您还可以为浅色和深色背景指定颜色选项:

```go
lipgloss.AdaptiveColor{Light: "236", Dark: "248"}
```

将自动检测终端的背景色，并运行时将选择适当的颜色。


## 内联格式

Lip Gloss支持常见的ANSI文本格式选项:

```go
var style = lipgloss.NewStyle().
    Bold(true).
    Italic(true).
    Faint(true).
    Blink(true).
    Strikethrough(true).
    Underline(true).
    Reverse(true)
```


## 块级格式

Lip Gloss 同样支持块级格式:

```go
// Padding
var style = lipgloss.NewStyle().
    PaddingTop(2).
    PaddingRight(4).
    PaddingBottom(2).
    PaddingLeft(4)

// Margins
var style = lipgloss.NewStyle().
    MarginTop(2).
    RightMarginRight(4).
    MarginBottom(2).
    MarginLeft(4)
```

页边距和填充也有速记语法，和CSS有着相同的语法格式:

```go
// 2 cells on all sides
lipgloss.NewStyle().Padding(2)

// 2 cells on the top and bottom, 4 cells on the left and right
lipgloss.NewStyle().Margin(2, 4)

// 1 cell on the top, 4 cells on the sides, 2 cells on the bottom
lipgloss.NewStyle().Padding(1, 4, 2)

// 顺时针方向，从顶部开始：顶部2个单元格，右侧4个单元格，右侧3个单元格
// 底部和左边各一个
lipgloss.NewStyle().Margin(2, 4, 3, 1)
```


## 文本对齐

可以将文本段落对齐到左侧、右侧或中间。

```go
var style = lipgloss.NewStyle().
    Width(24).
    Align(lipgloss.Left).  // align it left
    Align(lipgloss.Right). // no wait, align it right
    Align(lipgloss.Center) // just kidding, align it in the center
```


## 宽、高

设置最小宽度和高度是简单而直接的。

```go
var str = lipgloss.NewStyle().
    Width(24).
    Height(32).
    Foreground(lipgloss.Color("63")).
    Render("What’s for lunch?")
```


## 复制样式

使用 `Copy()`函数:

```go
var style = lipgloss.NewStyle().Foreground(lipgloss.Color("219"))

var wildStyle = style.Copy().Blink(true)
```

`Copy()` 对基础数据结构执行复制，确保一种样式的真实的、未引用的副本。不复制就有可能变异风格。


## 继承

样式可以从其他样式继承规则。继承时，只有未设置的规则在接收器上是继承的。

```go
var styleA = lipgloss.NewStyle().
    Foreground(lipgloss.Color("229")).
    Background(lipgloss.Color("63"))

//这里只继承背景色，因为前景
//颜色已经设置好了
var styleB = lipgloss.NewStyle().
    Foreground(lipgloss.Color("201")).
    Inherit(styleA)
```


## Unsetting Rules

All rules can be unset:

```go
var style = lipgloss.NewStyle().
    Bold(true).                        // make it bold
    UnsetBold().                       // jk don't make it bold
    Background(lipgloss.Color("227")). // yellow background
    UnsetBackground()                  // never mind
```

当一个规则被使用 `unset` 时，它将不会被继承或复制。


## 强制执行规则

有时，例如在开发组件时，您需要确定样式定义在UI中尊重其预期用途。这是`Inline`
`MaxWidth` 和 `MaxHeight`的作用：

```go
// 强制渲染到一行，忽略边距、填充和边框。
someStyle.Inline(true).Render("yadda yadda")

// 同时将渲染限制为五个单元格
someStyle.Inline(true).MaxWidth(5).Render("yadda yadda")

// 将渲染限制为5x5单元块
someStyle.MaxWidth(5).MaxHeight(5).Render("yadda yadda")
```

## 渲染

通常，您只需调用`Render(string)` 来渲染一个 `lipgloss.Style`:

```go
fmt.Println(lipgloss.NewStyle().Bold(true).Render("Hello, kitty."))
```

But you could also use the Stringer interface:

```go
var style = lipgloss.NewStyle().String("你好，猫咪。").Bold(true)

fmt.Printf("%s\n", style)
```


## 链接段落

还有一些用于水平和垂直连接的实用函数文本段落。

```go
// 沿底部边缘水平连接三个段落
lipgloss.HorizontalJoin(lipgloss.Bottom, paragraphA, paragraphB, paragraphC)

// 沿中心轴垂直连接两个段落
lipgloss.VerticalJoin(lipgloss.Center, paragraphA, paragraphB)

//横向连接三个段落，较短的段落对齐20%
//从顶部开始
lipgloss.HorizontalJoin(0.2, paragraphA, paragraphB, paragraphC)
```


## 在空白中放置文本

有时你只是想把一个文本块放在空白处。

```go
//在80个单元格宽的空间中水平居中放置段落。
//返回块的高度与输入段落一样。
block := lipgloss.PlaceHorizontal(80, lipgloss.Center, fancyStyledParagraph)

//在30格高的空格底部放一段。
//返回的文本块的宽度与输入段落一样。
block := lipgloss.PlaceVertical(30, lipgloss.Bottom, fancyStyledParagraph)

// 在30x80单元格的右下角放置段落。
block := lipgloss.Place(30, 80, lipgloss.Right, lipgloss.Bottom, fancyStyledParagraph)
```

您还可以设置空白的样式。有关详细信息： [the docs][docs].

***


## What about [Bubble Tea][tea]?

`Lip Gloss` 不能代替 `Bubble Tea`。相反，它是一种很好的 `Bubble Tea`同伴。它的目的是使组装终端用户界面视图尽可能简单有趣，这样您就可以专注于构建应用程序而不是关注你自己的低层次的布局细节。

简单地说，你可以用 `Lip Gloss` 来帮助建立你的 `Bubble Tea` 的界面。

[tea]: https://github.com/charmbracelet/tea


## Under the Hood

`Lip Gloss` 是建立在优秀的[Termenv][termenv]和[Reflow][reflow]的基础上的。
分别处理颜色和ANSI感知文本操作的库。
对于许多用例，`Termenv`和`Reflow`就足以满足您的需求。

[termenv]: https://github.com/muesli/termenv
[reflow]: https://github.com/muesli/reflow


## 渲染 Markdown

对于更多的以文档为中心的渲染解决方案，支持列表、表格和语法突出显示的代码可以查看[Glamour][glamour]，

基于样式表的标记呈现器。

[glamour]: https://github.com/charmbracelet/glamour


## License

[MIT](https://github.com/charmbracelet/lipgloss/raw/master/LICENSE)


***

Part of [Charm](https://charm.sh).

<a href="https://charm.sh/"><img alt="The Charm logo" src="https://stuff.charm.sh/charm-badge-unrounded.jpg" width="400"></a>

Charm热爱开源 • Charm loves open source


[docs]: https://pkg.go.dev/github.com/charmbracelet/lipgloss?tab=doc
