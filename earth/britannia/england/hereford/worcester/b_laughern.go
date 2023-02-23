package worcester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉恩LaughernBarony struct {
	feud.BaseBarony
}

var BLaughern拉恩 feud.Barony = &拉恩LaughernBarony{}

func init() {
	f := BLaughern拉恩.(*拉恩LaughernBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laughern",
		TitleName: "拉恩",
		TitleCode: "b_laughern",
	}
}
