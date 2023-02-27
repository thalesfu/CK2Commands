package zaozerye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克德拉KedraBarony struct {
	feud.BaseBarony
}

var BKedra克德拉 feud.Barony = &克德拉KedraBarony{}

func init() {
    f := BKedra克德拉.(*克德拉KedraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kedra",
		TitleName: "克德拉",
		TitleCode: "b_kedra",
	}
}
