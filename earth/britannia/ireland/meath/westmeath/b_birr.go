package westmeath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯尔BirrBarony struct {
	feud.BaseBarony
}

var BBirr伯尔 feud.Barony = &伯尔BirrBarony{}

func init() {
	f := BBirr伯尔.(*伯尔BirrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birr",
		TitleName: "伯尔",
		TitleCode: "b_birr",
	}
}
