package cyrenaica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆哈亚特MukhaytaBarony struct {
	feud.BaseBarony
}

var BMukhayta穆哈亚特 feud.Barony = &穆哈亚特MukhaytaBarony{}

func init() {
	f := BMukhayta穆哈亚特.(*穆哈亚特MukhaytaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mukhayta",
		TitleName: "穆哈亚特",
		TitleCode: "b_mukhayta",
	}
}
