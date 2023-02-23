package tihama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜斯BaysBarony struct {
	feud.BaseBarony
}

var BBays拜斯 feud.Barony = &拜斯BaysBarony{}

func init() {
	f := BBays拜斯.(*拜斯BaysBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bays",
		TitleName: "拜斯",
		TitleCode: "b_bays",
	}
}
