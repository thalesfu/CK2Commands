package abakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希拉ShiraBarony struct {
	feud.BaseBarony
}

var BShira希拉 feud.Barony = &希拉ShiraBarony{}

func init() {
    f := BShira希拉.(*希拉ShiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shira",
		TitleName: "希拉",
		TitleCode: "b_shira",
	}
}
