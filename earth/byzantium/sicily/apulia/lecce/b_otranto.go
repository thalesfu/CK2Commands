package lecce

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥特朗托OtrantoBarony struct {
	feud.BaseBarony
}

var BOtranto奥特朗托 feud.Barony = &奥特朗托OtrantoBarony{}

func init() {
    f := BOtranto奥特朗托.(*奥特朗托OtrantoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "otranto",
		TitleName: "奥特朗托",
		TitleCode: "b_otranto",
	}
}
