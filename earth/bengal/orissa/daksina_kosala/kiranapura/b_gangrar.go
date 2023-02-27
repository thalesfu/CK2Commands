package kiranapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乾罗GangrarBarony struct {
	feud.BaseBarony
}

var BGangrar乾罗 feud.Barony = &乾罗GangrarBarony{}

func init() {
    f := BGangrar乾罗.(*乾罗GangrarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gangrar",
		TitleName: "乾罗",
		TitleCode: "b_gangrar",
	}
}
