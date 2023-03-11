package terebovl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库切列门KuchelemynBarony struct {
	feud.BaseBarony
}

var BKuchelemyn库切列门 feud.Barony = &库切列门KuchelemynBarony{}

func init() {
    f := BKuchelemyn库切列门.(*库切列门KuchelemynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuchelemyn",
		TitleName: "库切列门",
		TitleCode: "b_kuchelemyn",
	}
}
