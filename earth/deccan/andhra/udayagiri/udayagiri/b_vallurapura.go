package udayagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伐露罗补罗VallurapuraBarony struct {
	feud.BaseBarony
}

var BVallurapura伐露罗补罗 feud.Barony = &伐露罗补罗VallurapuraBarony{}

func init() {
    f := BVallurapura伐露罗补罗.(*伐露罗补罗VallurapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vallurapura",
		TitleName: "伐露罗补罗",
		TitleCode: "b_vallurapura",
	}
}
