package kyzyl_su

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扬加贾YangadzhaBarony struct {
	feud.BaseBarony
}

var BYangadzha扬加贾 feud.Barony = &扬加贾YangadzhaBarony{}

func init() {
    f := BYangadzha扬加贾.(*扬加贾YangadzhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yangadzha",
		TitleName: "扬加贾",
		TitleCode: "b_yangadzha",
	}
}
