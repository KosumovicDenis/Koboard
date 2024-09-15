package soundboard

import (
	"errors"
	"fmt"
	"os"

	"github.com/KosumovicDenis/Koboard/internal/model"
	"google.golang.org/protobuf/proto"
)

func StartSoundboard() {
    data, err := os.ReadFile("soundboard.protobuf")
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("Initializing soundboard")
        initSoundboard()
    } else {
        chk(err)
    } 
    soundboard_data := model.Soundboard{}
    err = proto.Unmarshal(data, &soundboard_data)
}

func GetProfiles(m *[]*model.Profile) {
    data, err := os.ReadFile("soundboard.protobuf")  
    chk(err)
    temp_soundboard_data := model.Soundboard{}
    err = proto.Unmarshal(data, &temp_soundboard_data) 
    *m = temp_soundboard_data.Profiles
}

func initSoundboard() {
    soundboard_data := &model.Soundboard{
        Profiles: []*model.Profile{
            {
                Id: 1,
                Name: "Default",
                Active: true,
                Audios: []*model.Audio{
                },
            },
            {
                Id: 2,
                Name: "Secondo",
                Active: false,
                Audios: []*model.Audio{
                },
            },
            {
                Id: 3,
                Name: "Terzo",
                Active: false,
                Audios: []*model.Audio{
                },
            },
        },
    }
    data, err := proto.Marshal(soundboard_data)
    chk(err)
    chk(os.WriteFile("soundboard.protobuf", data, 0600))
}

func chk(err error) {
    if err != nil {
        panic(err)
    }
}
