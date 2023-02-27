package sunnmore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯克GiskeBarony struct {
	feud.BaseBarony
}

var BGiske伊斯克 feud.Barony = &伊斯克GiskeBarony{}

func init() {
    f := BGiske伊斯克.(*伊斯克GiskeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "giske",
		TitleName: "伊斯克",
		TitleCode: "b_giske",
	}
}
