package gegyai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 却藏ChuktsangBarony struct {
	feud.BaseBarony
}

var BChuktsang却藏 feud.Barony = &却藏ChuktsangBarony{}

func init() {
	f := BChuktsang却藏.(*却藏ChuktsangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chuktsang",
		TitleName: "却藏",
		TitleCode: "b_chuktsang",
	}
}
