package azov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库格伊KugeyBarony struct {
	feud.BaseBarony
}

var BKugey库格伊 feud.Barony = &库格伊KugeyBarony{}

func init() {
    f := BKugey库格伊.(*库格伊KugeyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kugey",
		TitleName: "库格伊",
		TitleCode: "b_kugey",
	}
}
