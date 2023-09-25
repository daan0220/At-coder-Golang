/*++++ a code submitted by daan0220 +++++++++++++
              /＼  ／ヽ
             {／￣￣￣ヽ!
             ∠＿＿╋＿＿ｊ
 AC ヨシ!   / (.)八(.)  ヽ
           ｛=/(人_)=|´￣)｀ヽ
            ＼ { {_,ﾉ ﾉ   //~ `
        ⊂￣ヽ_＞―――‐''’,〈   (＿)
         `ヘ(＿ ィ r―‐γ   !
              _,ﾉ j   |   |
            ｛   ｛    ﾉ  /＼
             ＼ス￣￣￣lしｲ＼ ＼
            (￣ ）     j /   ＼_j＼
             ￣￣     ( ｀ヽ   ＼__)
                       ＼__ﾉ
++++++++++++++++++++++++++++++++++++++++++++++++++*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

//+++++++++++++++++++++++++++++++++++++++
// 準備用の処理
//+++++++++++++++++++++++++++++++++++++++
var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

//+++++++++++++++++++++++++++++++++++++++
// main
//+++++++++++++++++++++++++++++++++++++++

func main() {
	defer func() { wr.Flush() }()

	var s string
	fmt.Scan(&s)

	pattern := "^[A-Z]([1-9][0-9]{5})[A-Z]$"
	match, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}
	if match {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

//+++++++++++++++++++++++++++++++++++++++
// 入力用の関数
//+++++++++++++++++++++++++++++++++++++++

//文字列を読み込む関数
func ins() string {
	sc.Scan()
	return sc.Text()
}

//Intを読み込む関数
func in() int {
	return atoi(ins())
}

// Intを読み込む関数
// 2個の変数にいっぺんで読み込むパターン
func in2() (int, int) {
	return atoi(ins()), atoi(ins())
}

// 浮動小数点数を読み込む関数
func infl() float64 {
	return atof(ins())
}

//+++++++++++++++++++++++++++++++++++++++
// 出力用の関数
//+++++++++++++++++++++++++++++++++++++++

//改行付き出力
func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

//フォーマット出力
func outf(s string, x ...interface{}) {
	fmt.Fprintf(wr, s, x...)
}

//+++++++++++++++++++++++++++++++++++++++
// 変換用の関数
//+++++++++++++++++++++++++++++++++++++++

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func atof(s string) float64 {
	f, e := strconv.ParseFloat(s, 64)
	if e != nil {
		panic(e)
	}
	return f
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
