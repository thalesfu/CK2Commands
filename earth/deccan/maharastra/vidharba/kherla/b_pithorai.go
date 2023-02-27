package kherla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗豆罗PithoraiBarony struct {
	feud.BaseBarony
}

var BPithorai毗豆罗 feud.Barony = &毗豆罗PithoraiBarony{}

func init() {
    f := BPithorai毗豆罗.(*毗豆罗PithoraiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pithorai",
		TitleName: "毗豆罗",
		TitleCode: "b_pithorai",
	}
}
