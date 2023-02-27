package magnitaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马格尼特纳亚MagnitayaBarony struct {
	feud.BaseBarony
}

var BMagnitaya马格尼特纳亚 feud.Barony = &马格尼特纳亚MagnitayaBarony{}

func init() {
    f := BMagnitaya马格尼特纳亚.(*马格尼特纳亚MagnitayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "magnitaya",
		TitleName: "马格尼特纳亚",
		TitleCode: "b_magnitaya",
	}
}
