package yarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 固玛GumaBarony struct {
	feud.BaseBarony
}

var BGuma固玛 feud.Barony = &固玛GumaBarony{}

func init() {
    f := BGuma固玛.(*固玛GumaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guma",
		TitleName: "固玛",
		TitleCode: "b_guma",
	}
}
