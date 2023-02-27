package bilino_polje

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷舍沃KresevoBarony struct {
	feud.BaseBarony
}

var BKresevo克雷舍沃 feud.Barony = &克雷舍沃KresevoBarony{}

func init() {
    f := BKresevo克雷舍沃.(*克雷舍沃KresevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kresevo",
		TitleName: "克雷舍沃",
		TitleCode: "b_kresevo",
	}
}
