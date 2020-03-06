package main

import "flag"
import "fmt"
// コマンドラインのコマンドと引数を取得する。
// 引数 なし
// 戻り ポインタ ログディレクトリ(Stringでも良いような気がするが)
//      String コマンド名
//      String[] コマンドパラメータ
func parseArgs()(string, string, []string){
	logDir := flag.String("logdir","","Log output directory (default=stderr)")
	// 引数：オプション名，初期値，ヘルプ文字列
	flag.Parse()
	fmt.Println("logDir", *logDir)
	fmt.Println("flag.Arg(0)", flag.Arg(0))
	fmt.Println("flag.Args", flag.Args()[1:])
	return *logDir, flag.Arg(0), flag.Args()[1:]
}
