package ghana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库加KughahBarony struct {
	feud.BaseBarony
}

var BKughah库加 feud.Barony = &库加KughahBarony{}

func init() {
    f := BKughah库加.(*库加KughahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kughah",
		TitleName: "库加",
		TitleCode: "b_kughah",
	}
}
