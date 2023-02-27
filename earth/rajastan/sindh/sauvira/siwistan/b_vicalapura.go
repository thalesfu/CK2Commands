package siwistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗遮罗补罗VicalapuraBarony struct {
	feud.BaseBarony
}

var BVicalapura毗遮罗补罗 feud.Barony = &毗遮罗补罗VicalapuraBarony{}

func init() {
    f := BVicalapura毗遮罗补罗.(*毗遮罗补罗VicalapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vicalapura",
		TitleName: "毗遮罗补罗",
		TitleCode: "b_vicalapura",
	}
}
