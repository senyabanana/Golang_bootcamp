package runner

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	// Перенаправляем вывод в буфер
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("echo", "hello")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	assert.NoError(t, err)
	assert.Equal(t, "hello\n", stdout.String())
	assert.Equal(t, "", stderr.String())
}

func TestProcessStdin(t *testing.T) {
	// Создаем pipe для os.Stdin
	stdinReader, stdinWriter, _ := os.Pipe()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = stdinReader

	// Создаем pipe для os.Stdout
	stdoutReader, stdoutWriter, _ := os.Pipe()
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	os.Stdout = stdoutWriter

	// Пишем в pipe, чтобы имитировать ввод пользователя
	input := "line2\nline1\n"
	io.WriteString(stdinWriter, input)
	stdinWriter.Close()

	// Запускаем ProcessStdin с командой echo
	err := ProcessStdin("echo", []string{})
	assert.NoError(t, err)

	// Закрываем stdoutWriter и читаем результат
	stdoutWriter.Close()
	var stdout bytes.Buffer
	io.Copy(&stdout, stdoutReader)

	// Проверяем, что вывод соответствует ожидаемому результату
	expectedOutput := "line1\nline2\n"
	actualOutput := stdout.String()
	assert.Equal(t, expectedOutput, actualOutput)
}
