package napata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尼塞KanisahBarony struct {
	feud.BaseBarony
}

var BKanisah凯尼塞 feud.Barony = &凯尼塞KanisahBarony{}

func init() {
    f := BKanisah凯尼塞.(*凯尼塞KanisahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanisah",
		TitleName: "凯尼塞",
		TitleCode: "b_kanisah",
	}
}
