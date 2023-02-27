package hebron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉特伦LatrunBarony struct {
	feud.BaseBarony
}

var BLatrun拉特伦 feud.Barony = &拉特伦LatrunBarony{}

func init() {
    f := BLatrun拉特伦.(*拉特伦LatrunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "latrun",
		TitleName: "拉特伦",
		TitleCode: "b_latrun",
	}
}
