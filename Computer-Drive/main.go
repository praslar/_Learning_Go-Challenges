package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := `my.song.mp3 11b 
greatSong.flac 1000b 
not3.txt 5b
video.mp4 200b 
game.exe 100b 
mov!e.mkv 10000b`
	fmt.Println(Solution(s))
}

func Solution(s string) string {

	files := strings.Split(s, "\n")
	sizeFile := make(map[string]int)

	types := map[string]string{
		"mp3":  "music",
		"aac":  "music",
		"flac": "music",

		"jpg": "images",
		"bmp": "images",
		"gif": "images",

		"mp4": "movies",
		"avi": "movies",
		"mkv": "movies",
	}

	for _, file := range files {
		f := strings.Split(file, " ")
		getChar := strings.Split(f[0], ".")
		ext := getChar[len(getChar)-1]
		bytes, _ := strconv.Atoi(string(f[1][:len(f[1])-1]))
		v, found := types[ext]
		if found {
			sizeFile[v] += bytes
		} else {
			sizeFile["other"] += bytes
		}
	}

	rs := fmt.Sprintf(`music %vb
images %vb
movies %vb
other %vb`, sizeFile["music"], sizeFile["images"], sizeFile["movies"], sizeFile["other"])
	return rs
}
