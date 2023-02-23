package qamdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 嘎玛KarmaBarony struct {
	feud.BaseBarony
}

var BKarma嘎玛 feud.Barony = &嘎玛KarmaBarony{}

func init() {
	f := BKarma嘎玛.(*嘎玛KarmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karma",
		TitleName: "嘎玛",
		TitleCode: "b_karma",
	}
}
