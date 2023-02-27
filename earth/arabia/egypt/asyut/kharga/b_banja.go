package kharga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班贾BanjaBarony struct {
	feud.BaseBarony
}

var BBanja班贾 feud.Barony = &班贾BanjaBarony{}

func init() {
    f := BBanja班贾.(*班贾BanjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banja",
		TitleName: "班贾",
		TitleCode: "b_banja",
	}
}
