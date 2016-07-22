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
    // counter
    X string // prev char key
    Y string // prev elem key
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
<title>The Return of Amontillado</title>
<meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
<meta name="viewport" content="width=1080,initial-scale=1" />
<link rel="stylesheet" type="text/css" href="https://fonts.googleapis.com/css?family=Amiko" />
<style type="text/css">
body{background-color:rgba(255,255,255,1.0);font-family:'Amiko',sans;}
p{padding:0px;font-size:18px;}
span.lines{color:rgba(35,35,35,1.0);background-color:rgba(230,230,230,1.0);margin:0px 40px 0px 6px;padding:4px}
span.script{white-space:pre;font-size:18px;}
span.character{padding:0px 0px 0px 160px;}
span.scene{font-weight:bold;}
span.paren{font-style:italic;padding:0px 0px 0px 100px;}
span.dialogue{padding:0px 0px 0px 100px;}
span.action{}
span.fadein{margin-left:400px;font-size:20px;}
span.fadeout{margin-left:400px;font-size:20px;}
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
        <br /><p>Contact: Gushi Tang &lt;gushitang@gmail.com&gt;</p>
    </body>
</html>`
    W.WriteString(ft)
}

func Body() {
    Meta()
    Character()
    Element()
    // Info()
    S.Scan()
    // S.Scan()
    S.Scan()
    i0 := 0
    for S.Scan() {
        // s0 := fmt.Sprintf("<p id=\"line_%d\">%d</p>\n", i0, i0)
        // W.WriteString(s0)
        s0 := S.Text()
        s1 := strings.Split(s0, " ")
        // s2 := s1[0]
        // fmt.Println(len(s2))
        // set curr and prev pointers
        // Y = X
        // X = s2
        // fmt.Printf("curr: %s, prev: %s\n", X, Y)
        // bold flag
        // b2 := false
        var s3 string
        if len(s1[0]) == 0 {
            // line break
            s3 = fmt.Sprintf("<p id=\"line_%d\"><span class=\"lines\">%d</span></p>\n", i0, i0)
        } else {
            // set curr and prev pointers
            // b0 := strings.Compare(s1[0], Y)
            // fmt.Printf("curr: %s, prev: %s, compare: %d\n", s1[0], Y, b0)
            // if DL or PA
            s4 := s1[1:]
            s5 := strings.Join(s4, " ")
            // fmt.Println(string(s5))
            if s1[0] == "DL" || s1[0] == "PA" {
                // compare char key
                b1 := strings.Compare(s1[1], X)
                if b1 != 0 {
                    // style margins for char
                    s6 := fmt.Sprintf("<p id=\"line_%d\"><span class=\"lines\">%d</span><span class=\"character\">%s</span></p>\n", i0, i0, C[s1[1]])
                    W.WriteString(s6)
                    i0 = i0 + 1
                }
                s4 = s1[2:]
                s5 = strings.Join(s4, " ")
                X = s1[1]
            }
            if s1[0] == "SC" || s1[0] == "AC" {
                X = ""
                // style bold for sc
                /*
                if s1[0] == "SC" {
                    b2 = true
                }
                */
            }
            /*
            if b2 == true {
                s3 = fmt.Sprintf("<p id=\"line_%d\"><span class=\"lines\">%d</span><span class=\"scene\">%s</span></p>\n", i0, i0, s5)
            } else {
                s3 = fmt.Sprintf("<p id=\"line_%d\"><span class=\"lines\">%d</span><span>%s</span></p>\n", i0, i0, s5)
            }
            */
            switch s1[0] {
                case "SC":
                    s3 = fmt.Sprintf("<p id=\"line_%d\"><span class=\"lines\">%d</span><span class=\"scene\">%s</span></p>\n", i0, i0, s5)
                case "DL":
                    s3 = fmt.Sprintf("<p id=\"line_%d\"><span class=\"lines\">%d</span><span class=\"dialogue\">%s</span></p>\n", i0, i0, s5)
                case "PA":
                    s3 = fmt.Sprintf("<p id=\"line_%d\"><span class=\"lines\">%d</span><span class=\"paren\">%s</span></p>\n", i0, i0, s5)
                case "AC":
                    s3 = fmt.Sprintf("<p id=\"line_%d\"><span class=\"lines\">%d</span><span class=\"action\">%s</span></p>\n", i0, i0, s5)
                case "FI":
                    s3 = fmt.Sprintf("<p id=\"line_%d\"><span class=\"lines\">%d</span><span class=\"fadein\">%s</span></p>\n", i0, i0, s5)
                case "FO":
                    s3 = fmt.Sprintf("<p id=\"line_%d\"><span class=\"lines\">%d</span><span class=\"fadeout\">%s</span></p>\n", i0, i0, s5)
            }
            Y = s1[0]
            // X = ""
        }
        // s2 := fmt.Sprintf("<p id=\"line_%d\">%d</p>\n", i0, s1[0][0])
        W.WriteString(s3)
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


