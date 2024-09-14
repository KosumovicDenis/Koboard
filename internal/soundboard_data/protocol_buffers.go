package soundboard_data

import (
	"os"

	"github.com/KosumovicDenis/Koboard/internal/model"
	"google.golang.org/protobuf/proto"
)

func InitSoundboard() {
    soundboard := &model.Soundboard{
        Profiles: []*model.Profile{
            {
                Id: 1,
                Name: "ProfileName",
                Audios: []*model.Audio{
                    {Id: 1, Name: "FileName", Path: "/path/to/audio/file"},
                },
            },
        },
    }
    data, err := proto.Marshal(soundboard)
    chk(err)
    chk(os.WriteFile("soundboard.protobuf", data, 0600))
}

func GetSoundBoard() {
    data, err := os.ReadFile("soundboard.protobuf")
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
