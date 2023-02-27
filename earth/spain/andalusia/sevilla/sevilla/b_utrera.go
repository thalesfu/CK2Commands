package sevilla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌特雷拉UtreraBarony struct {
	feud.BaseBarony
}

var BUtrera乌特雷拉 feud.Barony = &乌特雷拉UtreraBarony{}

func init() {
    f := BUtrera乌特雷拉.(*乌特雷拉UtreraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "utrera",
		TitleName: "乌特雷拉",
		TitleCode: "b_utrera",
	}
}
