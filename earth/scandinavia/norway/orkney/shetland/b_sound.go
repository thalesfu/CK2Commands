package shetland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑德SoundBarony struct {
	feud.BaseBarony
}

var BSound桑德 feud.Barony = &桑德SoundBarony{}

func init() {
    f := BSound桑德.(*桑德SoundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sound",
		TitleName: "桑德",
		TitleCode: "b_sound",
	}
}
