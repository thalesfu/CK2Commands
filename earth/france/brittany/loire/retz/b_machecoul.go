package retz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马什库勒MachecoulBarony struct {
	feud.BaseBarony
}

var BMachecoul马什库勒 feud.Barony = &马什库勒MachecoulBarony{}

func init() {
    f := BMachecoul马什库勒.(*马什库勒MachecoulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "machecoul",
		TitleName: "马什库勒",
		TitleCode: "b_machecoul",
	}
}
