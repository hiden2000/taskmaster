package cli

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/hiden2000/taskmaster/internal/storage"
)

func TestCLI(t *testing.T) {
	s := storage.NewStorage()
	cli := NewCLI(s)

	// 標準入出力をテストするためにPipeでI/Oを管理
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w // 出力を w (パイプの入り口)に設定

	// Test adding a task
	cli.executeCommand("add TestTask This is a test task")

	// Test listing tasks
	cli.executeCommand("list")

	// パイプの出力を確認
	w.Close()
	os.Stdout = oldStdout
	// バッファーに上記コマンドの実行結果をコピー
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Task added:") {
		t.Error("Expected 'Task added' message, but it was not found")
	}

	if !strings.Contains(output, "TestTask") {
		t.Error("Expected 'Test Task' in the list, but it was not found")
	}
}
