package makuria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯沃KawaBarony struct {
	feud.BaseBarony
}

var BKawa凯沃 feud.Barony = &凯沃KawaBarony{}

func init() {
    f := BKawa凯沃.(*凯沃KawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kawa",
		TitleName: "凯沃",
		TitleCode: "b_kawa",
	}
}
