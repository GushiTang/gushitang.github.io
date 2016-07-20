// html5 render for amontillado
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
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
    // maps
    M map[string]string // meta data
    C map[string]string // char map
    E map[string]string // elem map
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
<link rel="stylesheet" type="text/css" href="https://fonts.googleapis.com/css?family=Amiko" />
<style type="text/css">
body{background-color:rgba(255,255,255,1.0);font-family:'Amiko',sans;}
p{padding:2px;font-size:24px;}
span.lines{color:rgba(35,35,35,1.0);background-color:rgba(230,230,230,1.0);margin:0px 40px 0px 6px;padding:4px;text-alingn:center;}
span.script{white-space:pre;font-size:18px;)
</style>
<head>
</head>
<body>
`
    W.WriteString(hd)
    // google font: Amiko
}

func Foot() {
    ft := `<br /><p>Copyright &copy; 2016 Gushi Tang. All rights reserved.</p>
    </body>
</html>`
    W.WriteString(ft)
}

func Body() {
    Meta()
    Character()
    Element()
    i0 := 0
    for S.Scan() {
        s0 := fmt.Sprintf("<p id=\"line_%d\">%d</p>\n", i0, i0)
        W.WriteString(s0)
        i0 = i0 + 1
    }
}

func Meta() {
    // scan meta data
    M = make(map[string]string)
    for i := 0; i < 7; i++ {
        S.Scan()
        s0 := S.Text()
        // fmt.Println(s0)
        s1 := strings.Split(s0, " ")
        // fmt.Println(s1[:1],s1[1:])
        s2 := s1[:1][0]
        s3 := strings.Join(s1[1:], " ")
        // M[s1[:1][0]] = s1[1:]
        M[s2] = s3
    }
    fmt.Println(M)
}

func Character() {
    // scan char map
    C = make(map[string]string)
    S.Scan()
    S.Scan()
    for i := 0; i < 18; i++ {
        S.Scan()
        s0 := S.Text()
        s1 := strings.Split(s0, " ")
        s2 := s1[:1][0]
        s3 := strings.Join(s1[1:], " ")
        C[s2] = s3
    }
    fmt.Println(C)
}

func Element() {
    // scan elem map
    E = make(map[string]string)
    S.Scan()
    S.Scan()
    for i := 0; i < 6; i++ {
        S.Scan()
        s0 := S.Text()
        s1 := strings.Split(s0, " ")
        s2 := s1[:1][0]
        s3 := strings.Join(s1[1:], " ")
        E[s2] = s3
    }
    fmt.Println(E)
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
    Body()
    Foot()
    Clean()
}

// new file for html5 canvas 
// static renders
// digital bitmaps for amontillado logo


