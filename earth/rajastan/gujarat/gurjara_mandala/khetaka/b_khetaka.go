package khetaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 契吒迦KhetakaBarony struct {
	feud.BaseBarony
}

var BKhetaka契吒迦 feud.Barony = &契吒迦KhetakaBarony{}

func init() {
    f := BKhetaka契吒迦.(*契吒迦KhetakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khetaka",
		TitleName: "契吒迦",
		TitleCode: "b_khetaka",
	}
}
