package mzab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯马吉BermaneBarony struct {
	feud.BaseBarony
}

var BBermane伯马吉 feud.Barony = &伯马吉BermaneBarony{}

func init() {
	f := BBermane伯马吉.(*伯马吉BermaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bermane",
		TitleName: "伯马吉",
		TitleCode: "b_bermane",
	}
}
