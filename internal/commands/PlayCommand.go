package commands

import (
	"flag"
	"fmt"
	"github.com/Musik-Bot/Musik-Bot/internal/downloader"
	"github.com/Musik-Bot/Musik-Bot/internal/mysql"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
	"io/ioutil"
	"strings"
)

type DownloadStruct struct {
	ID       int    `json:"id"`
	Url      string `json:"url"`
	FileName string `json:"file_name"`
	Name     string `json:"name"`
}

func PlayCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	params := strings.Split(m.Content, " ")[1:]
	url := strings.Split(params[0], "&")[0]
	conn := mysql.GetConn()
	if strings.HasPrefix(url, "https://") {
		if !strings.HasPrefix(url, "https://www.youtube.com/watch?v=") && !strings.HasPrefix(url, "https://youtube.com/watch?v=") {
			emb := embed.NewEmbed()
			emb.Title = "Invalid youtube link"
			emb.Color = 0x1C76E8
			emb.Description = "Please choose an valid youtube link"
			s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
		} else {
			stmt, _ := conn.Prepare("SELECT * FROM `downloads` WHERE `url`=?")
			resp, err := stmt.Query(url)
			if err != nil {
				panic(err)
			}
			var downloads []DownloadStruct
			for resp.Next() {
				var cache DownloadStruct
				err = resp.Scan(&cache.ID, &cache.Url, &cache.FileName, &cache.Name)
				if err != nil {
					panic(err)
				}
				downloads = append(downloads, cache)
			}
			if len(downloads) == 0 {
				downloader.Download(url)
			} else {
				var (
					GuildID   = flag.String("g", "785392512880607253", "Guild ID")
					ChannelID = flag.String("c", "785392513313538120", "Channel ID")
					Folder    = flag.String("f", "./music", "Folder of files to play.")
					err       error
				)
				flag.Parse()
				dgv, err := s.ChannelVoiceJoin(*GuildID, *ChannelID, false, true)
				if err != nil {
					fmt.Println(err)
					return
				}

				// Start loop and attempt to play all files in the given folder
				files, _ := ioutil.ReadDir(*Folder)
				for _, f := range files {
					fmt.Println("PlayAudioFile:", f.Name())
					s.UpdateStatus(0, f.Name())

					dgvoice.PlayAudioFile(dgv, fmt.Sprintf("%s/%s", *Folder, f.Name()), make(chan bool))
				}

				// Close connections
				dgv.Close()
				s.Close()

			}
		}
	} else {
		fmt.Println("SearchValue")
	}
}
