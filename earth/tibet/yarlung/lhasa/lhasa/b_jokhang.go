package lhasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 觉康JokhangBarony struct {
	feud.BaseBarony
}

var BJokhang觉康 feud.Barony = &觉康JokhangBarony{}

func init() {
    f := BJokhang觉康.(*觉康JokhangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jokhang",
		TitleName: "觉康",
		TitleCode: "b_jokhang",
	}
}
