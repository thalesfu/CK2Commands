package cakrakuta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗耶姞利呬RaigarhBarony struct {
	feud.BaseBarony
}

var BRaigarh罗耶姞利呬 feud.Barony = &罗耶姞利呬RaigarhBarony{}

func init() {
    f := BRaigarh罗耶姞利呬.(*罗耶姞利呬RaigarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raigarh",
		TitleName: "罗耶姞利呬",
		TitleCode: "b_raigarh",
	}
}
