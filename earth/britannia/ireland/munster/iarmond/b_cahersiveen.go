package iarmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔西温CahersiveenBarony struct {
	feud.BaseBarony
}

var BCahersiveen凯尔西温 feud.Barony = &凯尔西温CahersiveenBarony{}

func init() {
	f := BCahersiveen凯尔西温.(*凯尔西温CahersiveenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cahersiveen",
		TitleName: "凯尔西温",
		TitleCode: "b_cahersiveen",
	}
}
