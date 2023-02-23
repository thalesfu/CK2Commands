package wadai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦拉WaraBarony struct {
	feud.BaseBarony
}

var BWara瓦拉 feud.Barony = &瓦拉WaraBarony{}

func init() {
	f := BWara瓦拉.(*瓦拉WaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wara",
		TitleName: "瓦拉",
		TitleCode: "b_wara",
	}
}
