package soundboard_data

import (
	"os"
	"strings"

	"github.com/KosumovicDenis/Koboard/internal/model"
	"google.golang.org/protobuf/proto"
)

func NewSoundBoard(name string) {
    soundboard := &model.Soundboard{
        Id: 1,
        Name: name,
        Audios: []*model.Audio{
            {Id: 1, Name: "Nome file", Path: "Path file"},
        },
    }
    data, err := proto.Marshal(soundboard)
    chk(err)
    filename := "soundboard_" + strings.ToLower(strings.ReplaceAll(name, " ", "_" )) + ".protobuf"
    chk(os.WriteFile(filename, data, 0600))
}

func GetSoundBoard(name string) {
    filename := "soundboard_" + strings.ToLower(strings.ReplaceAll(name, " ", "_" )) + ".protobuf"
    data, err := os.ReadFile(filename)
    chk(err)
    soundboard := model.Soundboard{}
    err = proto.Unmarshal(data, &soundboard)
    chk(err)
}

func chk(err error) {
    if err != nil {
        panic(err)
    }
}
