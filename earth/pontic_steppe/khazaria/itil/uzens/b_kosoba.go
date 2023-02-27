package uzens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科索巴KosobaBarony struct {
	feud.BaseBarony
}

var BKosoba科索巴 feud.Barony = &科索巴KosobaBarony{}

func init() {
    f := BKosoba科索巴.(*科索巴KosobaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kosoba",
		TitleName: "科索巴",
		TitleCode: "b_kosoba",
	}
}
