package main

import "io"
import "os"
import "os/exec"

func execution(commandName string,
               args[] string,
			   stdout io.Writer,
			   stderr io.Writer) (*os.ProcessState, error){
	// 子プロセスの準備
	cmd := exec.Command(commandName, args...)

	// 子プロセスの準備 標準出力を取り出す
	childStdout, _ := cmd.StdoutPipe()
	childStderr, _ := cmd.StderrPipe()

	// 子プロセスの準備 取り出した標準出力を並行処理でio.Writerに書出す
	go io.Copy(stdout, childStdout)
	go io.Copy(stderr, childStderr)

	// 子プロセス実行
	err := cmd.Run()
	if err 	!= nil {
		return nil, err
	}
	return cmd.ProcessState, nil
}

