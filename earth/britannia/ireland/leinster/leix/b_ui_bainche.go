package leix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊巴尔赫Ui_baincheBarony struct {
	feud.BaseBarony
}

var BUi_bainche伊巴尔赫 feud.Barony = &伊巴尔赫Ui_baincheBarony{}

func init() {
    f := BUi_bainche伊巴尔赫.(*伊巴尔赫Ui_baincheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ui_bainche",
		TitleName: "伊巴尔赫",
		TitleCode: "b_ui_bainche",
	}
}
