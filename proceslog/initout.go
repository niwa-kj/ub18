package main

import "fmt"
import "io"
import "os"
import "path/filepath"
import "time"
// ログファイルを初期化する
// 引数 String ログディレクトリ
//      String コマンド名
// 戻り io.Writer 標準出力
//      io.Writer 標準エラー出力
//      error エラーインタフェース(正常時:nil)
//横に伸びない書き方はないものだろうか？
func initOut(logDir string,
             commandName string) (stdout io.Writer,
			                      stderr io.Writer,
								  err error){
	if logDir == "" {
		stdout = os.Stdout
		stderr = os.Stderr
	}else{
		// ファイル名:コマンド名-タイムスタンプ-stdout.log
		ts := time.Now().Unix()
		stdoutFileName := fmt.Sprintf("%s-%v-stdout.log", commandName, ts)
		sdtoutFile, err := os.Create(filepath.Join(logDir,stdoutFileName))
		if err != nil{
			return nil, nil, err
		}
		stdout = io.MultiWriter(os.Stdout, sdtoutFile)


		stderrFileName := fmt.Sprintf("%s-%v-stderr.log", commandName, ts)
		sdterrFile, err := os.Create(filepath.Join(logDir,stderrFileName))
		if err != nil{
			return nil, nil, err
		}

		stderr = io.MultiWriter(os.Stderr, sdterrFile)
	}
	return
}


