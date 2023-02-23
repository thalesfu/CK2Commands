package tobol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊西利库尔IsilkulBarony struct {
	feud.BaseBarony
}

var BIsilkul伊西利库尔 feud.Barony = &伊西利库尔IsilkulBarony{}

func init() {
	f := BIsilkul伊西利库尔.(*伊西利库尔IsilkulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isilkul",
		TitleName: "伊西利库尔",
		TitleCode: "b_isilkul",
	}
}
