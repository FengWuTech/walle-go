package waller

import (
	"fmt"
	"github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"github.com/go-cmd/cmd"
	socketio "github.com/googollee/go-socket.io"
	"golang.org/x/crypto/ssh"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
	"walle-go/logger"
	"walle-go/models"
	"walle-go/models/record_model"
)

type Waller struct {
	Dir    string
	Env    []string
	Host   string
	Port   int
	User   string
	Room   string
	Cmd    *cmd.Cmd
	Server *socketio.Server
}

func (w *Waller) LoadEnv(envs map[string]string) error {
	// 清空
	w.Env = []string{}
	for k, v := range envs {
		w.Env = append(w.Env, fmt.Sprintf(`%s=%s`, k, v))
	}
	return nil
}

func (w *Waller) Run(wenv map[string]string, command string) error {
	// Disable output buffering, enable streaming
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}
	stdoutLines := make([]string, 0)
	stderrLines := make([]string, 0)

	// Create Cmd with options
	runCmd := cmd.NewCmdOptions(cmdOptions, "bash", "-c", command)
	runCmd.Env = w.Env
	runCmd.Dir = w.Dir
	w.Cmd = runCmd
	logger.Debugf("Run Command %s at Dir %s with Envs: %v", command, runCmd.Dir, runCmd.Env)

	// Print STDOUT and STDERR lines streaming from Cmd
	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)
		// Done when both channels have been closed
		// https://dave.cheney.net/2013/04/30/curious-channels
		for runCmd.Stdout != nil || runCmd.Stderr != nil {
			select {
			case line, open := <-runCmd.Stdout:
				if !open {
					runCmd.Stdout = nil
					continue
				}
				logger.Debugf("Run Command Output: %s", line)
				stdoutLines = append(stdoutLines, line)
			case line, open := <-runCmd.Stderr:
				if !open {
					runCmd.Stderr = nil
					continue
				}
				logger.Debugf("Run Command stderr Output: %s", line)
				stderrLines = append(stderrLines, line)
			}
		}
	}()

	// Stop command after 10 minutes
	go func() {
		<-time.After(10 * time.Minute)
		runCmd.Stop()
	}()

	statusChan := runCmd.Start() // non-blocking
	// Wait for goroutine to print everything
	<-doneChan
	// Block waiting for command to exit, be stopped, or be killed
	<-statusChan
	status := runCmd.Status()
	stage := wenv["stage"]
	taskId, _ := strconv.Atoi(wenv["task_id"])
	sequence, _ := strconv.Atoi(wenv["sequence"])
	userId, _ := strconv.Atoi(wenv["user_id"])
	stdoutOneLine := strings.Join(stdoutLines, "\n")
	stderrOneLine := strings.Join(stderrLines, "\n")
	exitCode := 0
	if status.Error != nil {
		exitCode = status.Exit
	}
	wsDict := map[string]interface{}{
		"user_model": w.User,
		"host": w.Host,
		"cmd": command,
		"status": exitCode,
		"stage": stage,
		"sequence": sequence,
		"success": stdoutOneLine,
		"error": stderrOneLine,
	}
	// 发送消息 TODO 封一个Emit方法
	w.Server.BroadcastToRoom("/walle", w.Room, "console", map[string]interface{}{"event": "task:console", "data": wsDict})
	record := record_model.Records{
		Stage:     &stage,
		Sequence:  &sequence,
		UserId:    &userId,
		TaskId:    &taskId,
		Status:    &exitCode,
		Host:      &w.Host,
		User:      &w.User,
		Command:   &command,
		Success:   &stdoutOneLine,
		Error:     &stderrOneLine,
	}
	models.GetDB().Create(&record)
	if status.Error != nil {
		return status.Error
	}
	return nil
}

func (w *Waller) Put(wenv map[string]string, local string, remote string) error {
	// Use SSH key authentication from the auth package
	// we ignore the host key in this example, please change this if you use this library
	currentUser, _ := user.Current()
	// clientConfig, _ := auth.PrivateKey("work", "/home/walle/.ssh/id_rsa", ssh.InsecureIgnoreHostKey())
	clientConfig, _ := auth.PrivateKey(w.User, currentUser.HomeDir+"/.ssh/id_rsa", ssh.InsecureIgnoreHostKey())
	// clientConfig, _ := auth.SshAgent(w.User, ssh.InsecureIgnoreHostKey())

	// For other authentication methods see ssh.ClientConfig and ssh.AuthMethod
	// clientConfig.Auth = append(clientConfig.Auth, ssh.Password(""))

	// Create a new SCP client
	client := scp.NewClient(fmt.Sprintf("%s:%d", w.Host, w.Port), &clientConfig)

	// Connect to the remote server
	err := client.Connect()
	if err != nil {
		logger.Error("Couldn't establish a connection to the remote server ", err)
		return err
	}

	// Open a file
	f, _ := os.Open(local)

	// Close client connection after the file has been copied
	defer client.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	err = client.CopyFile(f, remote, "0655")

	if err != nil {
		logger.Error("Error while copying file ", err)
		return err
	}
	logger.Debugf("scp from %s to %s:%s", local, w.Host, remote)
	stage := wenv["stage"]
	taskId, _ := strconv.Atoi(wenv["task_id"])
	sequence, _ := strconv.Atoi(wenv["sequence"])
	userId, _ := strconv.Atoi(wenv["user_id"])
	cmdStr := fmt.Sprintf("scp %s %s:%d%s", local, w.Host, w.Port, remote)
	exitCode := 0

	wsDict := map[string]interface{}{
		"user_model": w.User,
		"host": w.Host,
		"cmd": cmdStr,
		"status": exitCode,
		"stage": stage,
		"sequence": sequence,
		"success": "",
		"error": "",
	}
	// 发送消息
	w.Server.BroadcastToRoom("/walle", w.Room, "console", map[string]interface{}{"event": "task:console", "data": wsDict})

	record := record_model.Records{
		Stage:     &stage,
		Sequence:  &sequence,
		UserId:    &userId,
		TaskId:    &taskId,
		Status:    &exitCode,
		Host:      &w.Host,
		User:      &w.User,
		Command:   &cmdStr,
	}
	models.GetDB().Create(&record)
	return nil
}

func (w *Waller) RemoteRun(wenv map[string]string, command string) (string, error) {
	currentUser, _ := user.Current()
	clientConfig, _ := auth.PrivateKey(w.User, currentUser.HomeDir+"/.ssh/id_rsa", ssh.InsecureIgnoreHostKey())

	// For other authentication methods see ssh.ClientConfig and ssh.AuthMethod
	// clientConfig.Auth = append(clientConfig.Auth, ssh.Password(""))

	// Create a new SCP client
	client := scp.NewClient(fmt.Sprintf("%s:%d", w.Host, w.Port), &clientConfig)

	// Connect to the remote server
	err := client.Connect()
	if err != nil {
		logger.Error("Couldn't establish a connection to the remote server ", err)
		return "", err
	}

	output, err := client.Session.CombinedOutput(command)
	outputStr := string(output)
	logger.Debugf("Run Remote Command Output: %s", outputStr)
	stage := wenv["stage"]
	taskId, _ := strconv.Atoi(wenv["task_id"])
	sequence, _ := strconv.Atoi(wenv["sequence"])
	userId, _ := strconv.Atoi(wenv["user_id"])
	exitCode := 0

	wsDict := map[string]interface{}{
		"user_model": w.User,
		"host": w.Host,
		"cmd": command,
		"status": exitCode,
		"stage": stage,
		"sequence": sequence,
		"success": outputStr,
		"error": "",
	}
	// 发送消息
	w.Server.BroadcastToRoom("/walle", w.Room, "console", map[string]interface{}{"event": "task:console", "data": wsDict})

	record := record_model.Records{
		Stage:     &stage,
		Sequence:  &sequence,
		UserId:    &userId,
		TaskId:    &taskId,
		Status:    &exitCode,
		Host:      &w.Host,
		User:      &w.User,
		Command:   &command,
		Success:   &outputStr,
	}
	models.GetDB().Create(&record)
	return outputStr, nil
}
