package downloader

import (
	"fmt"
	"github.com/Musik-Bot/Musik-Bot/internal/mysql"
	"os/exec"
	"strings"
)

func Download(url string) {
	out, err2 := exec.Command("youtube-dl", "--extract-audio", "--audio-format", "mp3", "--output", "./music/%(title)s.%(ext)s", "--restrict-filenames", url).Output()
	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		fmt.Println("downloading file:")
		fmt.Println(string(out))
		name, _ := exec.Command("youtube-dl", "--extract-audio", "--audio-format", "mp3", "--output", "./music/%(title)s.%(ext)s", "--restrict-filenames", "--get-filename", url).Output()
		fname := strings.ReplaceAll(string(name), ".webm", ".mp3")
		conn := mysql.GetConn()
		stmt, _ := conn.Prepare("INSERT INTO `downloads` (`ID`, `url`, `FileName`, `Name`) VALUES (NULL, ?,?,?);")
		stmt.Exec(url, fname, strings.Split(fname, ".mp3")[0])
		defer conn.Close()
	}
}
