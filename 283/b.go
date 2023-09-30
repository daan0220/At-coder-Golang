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
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var t int
	fmt.Fscan(rd, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(rd, &n)

		var a, ans int
		for j := 0; j < n; j++ {
			fmt.Fscan(rd, &a)
			if a%2 == 1 {
				ans++
			}
		}

		fmt.Fprintln(wr, ans)
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