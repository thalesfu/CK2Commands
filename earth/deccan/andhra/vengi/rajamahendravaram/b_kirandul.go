package rajamahendravaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基兰杜尔KirandulBarony struct {
	feud.BaseBarony
}

var BKirandul基兰杜尔 feud.Barony = &基兰杜尔KirandulBarony{}

func init() {
	f := BKirandul基兰杜尔.(*基兰杜尔KirandulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirandul",
		TitleName: "基兰杜尔",
		TitleCode: "b_kirandul",
	}
}
