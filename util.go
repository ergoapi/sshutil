package sshutil

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func (ss *SSH) log(msg string)  {
	if ss.Debug {
		log.Println(msg)
	}
}

func Md5FromLocal(localPath string) string {
	cmd := fmt.Sprintf("md5sum %s | cut -d\" \" -f1", localPath)
	c := exec.Command("sh", "-c", cmd)
	out, err := c.CombinedOutput()
	if err != nil {
		return ""
	}
	md5 := string(out)
	md5 = strings.ReplaceAll(md5, "\n", "")
	md5 = strings.ReplaceAll(md5, "\r", "")

	return md5
}