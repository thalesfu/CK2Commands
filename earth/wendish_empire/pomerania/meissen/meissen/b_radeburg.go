package meissen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉德堡RadeburgBarony struct {
	feud.BaseBarony
}

var BRadeburg拉德堡 feud.Barony = &拉德堡RadeburgBarony{}

func init() {
    f := BRadeburg拉德堡.(*拉德堡RadeburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radeburg",
		TitleName: "拉德堡",
		TitleCode: "b_radeburg",
	}
}
