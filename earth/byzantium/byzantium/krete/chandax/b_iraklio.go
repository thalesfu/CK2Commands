package chandax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊拉克利翁IraklioBarony struct {
	feud.BaseBarony
}

var BIraklio伊拉克利翁 feud.Barony = &伊拉克利翁IraklioBarony{}

func init() {
	f := BIraklio伊拉克利翁.(*伊拉克利翁IraklioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iraklio",
		TitleName: "伊拉克利翁",
		TitleCode: "b_iraklio",
	}
}
