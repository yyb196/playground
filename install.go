package main

import (
	"os/exec"
	"fmt"
	"bytes"
	"strings"
	"os"
	"path/filepath"
)

const (
	Svn = "http://3xcd3fc59624a1.repo1.svn.jae.taobao.com/trunk"
)

func install(lang string) {
	//校验java有没有装，以及是否是1.7+的版本
	checkJavaVersion := exec.Command("java", "-version")
	output, err := GetOutput(checkJavaVersion)
	if err != nil {
		fmt.Println("failed.", err)
	}
	java7Flag := "java version \"1.7"
	java8Flag := "java version \"1.8"
	if !strings.Contains(output, java7Flag) && !strings.Contains(output, java8Flag) {
		fmt.Println("安装错误! 请安装1.7+的java版本并设置JAVA_HOME后执行本脚本!")
		return;
	}

	//获取操作系统信息
	checkUname := exec.Command("uname")
	output, err = GetOutput(checkUname)
	if err != nil {
		fmt.Println("failed.", err)
	}
	//	fmt.Fprintln(os.Stderr, output)

	//获取安装路径
	HOME := os.Getenv("HOME")

	defaultDir := HOME + "/TAE_Cloud_SDK"

	var sdkDir string
	fmt.Print("请输入sdk的安装目录[" + defaultDir + "]:")
	fmt.Scanln(&sdkDir)

	if len(strings.TrimSpace(sdkDir)) == 0 {
		sdkDir = defaultDir
	}
	fmt.Printf("开始安装 tae sdk 到:%s\n", sdkDir)
	currentPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	os.Setenv("user.dir", currentPath);

	Output(exec.Command("java", "-jar", currentPath+"/lib/svnclient.jar", lang, sdkDir, Svn, currentPath))
	/*
	Output(exec.Command("chmod", "755", sdkDir + "/tae-dev/bin/startup.sh"))
	Output(exec.Command("chmod", "755", sdkDir + "/tae-dev/bin/start.sh"))
	Output(exec.Command("chmod", "755", sdkDir + "/tae-dev/bin/catalina.sh"))
	Output(exec.Command("chmod", "755", sdkDir + "/tae-dev/bin/start.sh"))
	Output(exec.Command("chmod", "755", sdkDir + "/tae-dev/bin/restart.sh"))
	Output(exec.Command("chmod", "755", sdkDir + "/tae-dev/bin/restartup.sh"))
	Output(exec.Command("chmod", "755", sdkDir + "/tae-dev/bin/shutdown.sh"))
	fmt.Println("开始启动SDK......")
	Output(exec.Command(sdkDir + "/tae-dev/bin/startup.sh"))
	*/
	//	fmt.Println("启动成功，以后可以直接执行 " + sdkDir + "/tae-dev/bin/startup.sh 来启动SDK")
}

func GetOutput(cmd *exec.Cmd) (string, error) {
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func Output(cmd *exec.Cmd) {
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func main(){
	install("Java")
}

