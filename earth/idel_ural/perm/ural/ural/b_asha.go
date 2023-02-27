package ural

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿沙AshaBarony struct {
	feud.BaseBarony
}

var BAsha阿沙 feud.Barony = &阿沙AshaBarony{}

func init() {
    f := BAsha阿沙.(*阿沙AshaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asha",
		TitleName: "阿沙",
		TitleCode: "b_asha",
	}
}
