package daylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赞詹ZanjanBarony struct {
	feud.BaseBarony
}

var BZanjan赞詹 feud.Barony = &赞詹ZanjanBarony{}

func init() {
    f := BZanjan赞詹.(*赞詹ZanjanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zanjan",
		TitleName: "赞詹",
		TitleCode: "b_zanjan",
	}
}
