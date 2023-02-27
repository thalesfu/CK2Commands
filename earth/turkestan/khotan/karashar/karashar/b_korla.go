package karashar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坤闾KorlaBarony struct {
	feud.BaseBarony
}

var BKorla坤闾 feud.Barony = &坤闾KorlaBarony{}

func init() {
    f := BKorla坤闾.(*坤闾KorlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korla",
		TitleName: "坤闾",
		TitleCode: "b_korla",
	}
}
