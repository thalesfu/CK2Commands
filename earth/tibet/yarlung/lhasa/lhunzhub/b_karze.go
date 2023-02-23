package lhunzhub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡孜KarzeBarony struct {
	feud.BaseBarony
}

var BKarze卡孜 feud.Barony = &卡孜KarzeBarony{}

func init() {
	f := BKarze卡孜.(*卡孜KarzeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karze",
		TitleName: "卡孜",
		TitleCode: "b_karze",
	}
}
