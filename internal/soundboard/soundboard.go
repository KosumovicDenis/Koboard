package soundboard

import (
	"errors"
	"fmt"
	"os"

	"github.com/KosumovicDenis/Koboard/internal/model"
	"google.golang.org/protobuf/proto"
)

var soundboard_data model.Soundboard

func StartSoundboard() {
    data, err := os.ReadFile("soundboard.protobuf")
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("Initializing soundboard")
        initSoundboard()
    } else {
        chk(err)
        err = proto.Unmarshal(data, &soundboard_data)
    } 
}

func GetProfiles(m *[]*model.Profile) {
    *m = soundboard_data.Profiles
}

func initSoundboard() {
    soundboard_data = model.Soundboard{
        Profiles: []*model.Profile{
            {
                Id: 1,
                Name: "Default",
                Active: true,
                Audios: []*model.Audio{
                },
            },
        },
    }
    data, err := proto.Marshal(&soundboard_data)
    chk(err)
    chk(os.WriteFile("soundboard.protobuf", data, 0600))
}

func SaveSoundboard() {
    data, err := proto.Marshal(&soundboard_data)
    chk(err)
    chk(os.WriteFile("soundboard.protobuf", data, 0600))
}

func chk(err error) {
    if err != nil {
        panic(err)
    }
}
