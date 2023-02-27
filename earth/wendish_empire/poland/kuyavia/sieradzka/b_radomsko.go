package sieradzka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉多姆斯科RadomskoBarony struct {
	feud.BaseBarony
}

var BRadomsko拉多姆斯科 feud.Barony = &拉多姆斯科RadomskoBarony{}

func init() {
    f := BRadomsko拉多姆斯科.(*拉多姆斯科RadomskoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radomsko",
		TitleName: "拉多姆斯科",
		TitleCode: "b_radomsko",
	}
}
