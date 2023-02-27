package sennar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴卡赫BakkahBarony struct {
	feud.BaseBarony
}

var BBakkah巴卡赫 feud.Barony = &巴卡赫BakkahBarony{}

func init() {
    f := BBakkah巴卡赫.(*巴卡赫BakkahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakkah",
		TitleName: "巴卡赫",
		TitleCode: "b_bakkah",
	}
}
