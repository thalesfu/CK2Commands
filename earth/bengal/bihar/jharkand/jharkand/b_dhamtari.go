package jharkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀姆多利DhamtariBarony struct {
	feud.BaseBarony
}

var BDhamtari陀姆多利 feud.Barony = &陀姆多利DhamtariBarony{}

func init() {
    f := BDhamtari陀姆多利.(*陀姆多利DhamtariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhamtari",
		TitleName: "陀姆多利",
		TitleCode: "b_dhamtari",
	}
}
