// scan and render html5 for monk4.txt
package main

import (
    "bufio"
    "fmt"
    "html"
    "os"
    // "strings"
)

const (
    INPUT = "monk4.txt"
    OUTPUT = "apmon.html"
)

var (
    Apin *os.File
    Apout *os.File
    Scnr *bufio.Scanner
    Wrtr *bufio.Writer
)

func Load() {
    var err error
    Apin, err = os.Open(INPUT)
    if err != nil {
        fmt.Println(err)
    }
    Apout, err = os.Create(OUTPUT)
    if err != nil {
        fmt.Println(err)
    }
}

func Bufs() {
    Scnr = bufio.NewScanner(Apin)
    Wrtr = bufio.NewWriter(Apout)
    // fmt.Println(Scnr, Wrtr)
}

func Head() {
    // write head html5 string
    hd := `<!DOCTYPE>
<html>
<title>Apocalypse Mountain by Gushi Tang &lt;gushitang@gmail.com&gt;</title>
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
// hi gushi apmon!
// <br />
// all lato gato
// `
    // fmt.Println(hd)
    Wrtr.WriteString(hd)
}

func Foot() {
    ft := `<br /><p>Copyright &copy; 2016 Gushi Tang. All rights reserved.</p>
    </body>
</html>`
    // fmt.Println(ft)
    Wrtr.WriteString(ft)
}

func Body() {
    i0 := 0
    for Scnr.Scan() {
        i0 = i0 + 1
        s0 := Scnr.Text()
        s1 := html.EscapeString(s0)
        // s2 := strings.TrimSpace(s1)
        // fmt.Println(s2)
        s2 := fmt.Sprintf("<p><span id=\"lineno_%d\" class=\"lines\">%d</span><span id=\"script_%d\" class=\"script\">%s</code></p>\n", i0, i0, i0, s1)
        Wrtr.WriteString(s2)
    }
}

func Clean() {
    defer Apin.Close()
    defer Apout.Close()
    fmt.Println("closing open files")
    Wrtr.Flush()
}

func main() {
    fmt.Println("starting monk program :-)")
    fmt.Printf("input file: %s and output: %s.\n", INPUT, OUTPUT)
    fmt.Println("begin loading")
    Load()
    Bufs()
    Head()
    Body()
    Foot()
    Clean()
    fmt.Println("finished.")
}

// use p, header tags for heading elements
// for each line, use code for line nums
// and then just display text
// pre, margin or &nbsp; for tab white spacing


