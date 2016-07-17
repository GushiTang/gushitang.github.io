// html5 render for amontillado
package main

import (
    "bufio"
    "fmt"
    "os"
)

const (
    INPUT = "return_21.txt"
    OUTPUT = "return.html"
)

var (
    Retin *os.File
    Retout *os.File
    err error
    S *bufio.Scanner
    W *bufio.Writer
)

func Load() {
    Retin, err = os.Open(INPUT)
    if err != nil {
        fmt.Println(err)
    }
    Retout, err = os.Create(OUTPUT)
    if err != nil {
        fmt.Println(err)
    }
}

func Bufs() {
    S = bufio.NewScanner(Retin)
    W = bufio.NewWriter(Retout)
}

func Head() {
    // write head html5 string
    hd := `<!DOCTYPE>
<html>
<title>The Return of Amontillado by Gushi Tang &lt;gushitang@gmail.com&gt;</title>
<meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
<meta name="viewport" content="width=1080,initial-scale=1" />
<link rel="stylesheet" type="text/css" href="https://fonts.googleapis.com/css?family=Lato" />
<style type="text/css">
body{background-color:rgba(255,255,255,1.0);font-family:'Lato',sans;}
p{padding:2px;font-size:14px;}
span.lines{color:rgba(35,35,35,1.0);background-color:rgba(230,230,230,1.0);margin:0px 40px 0px 6px;padding:4px;text-alingn:center;}
span.script{white-space:pre;font-size:18px;)
</style>
<head>
</head>
<body>
`
    W.WriteString(hd)
}

func Foot() {
    ft := `<br /><p>Copyright &copy; 2016 Gushi Tang. All rights reserved.</p>
    </body>
</html>`
    W.WriteString(ft)
}


func Clean() {
    defer Retin.Close()
    defer Retout.Close()
    W.Flush()
}

func main() {
    fmt.Println("viva fortunado!")
    Load()
    Bufs()
    Head()
    Foot()
    Clean()
}
