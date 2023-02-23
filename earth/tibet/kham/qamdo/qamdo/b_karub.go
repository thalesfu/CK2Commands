package qamdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡若KarubBarony struct {
	feud.BaseBarony
}

var BKarub卡若 feud.Barony = &卡若KarubBarony{}

func init() {
	f := BKarub卡若.(*卡若KarubBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karub",
		TitleName: "卡若",
		TitleCode: "b_karub",
	}
}
