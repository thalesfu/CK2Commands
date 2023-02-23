package irgiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔吉兹IrgizBarony struct {
	feud.BaseBarony
}

var BIrgiz伊尔吉兹 feud.Barony = &伊尔吉兹IrgizBarony{}

func init() {
	f := BIrgiz伊尔吉兹.(*伊尔吉兹IrgizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "irgiz",
		TitleName: "伊尔吉兹",
		TitleCode: "b_irgiz",
	}
}
