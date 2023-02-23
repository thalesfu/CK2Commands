package vasyugan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宰采瓦ZaytsevaBarony struct {
	feud.BaseBarony
}

var BZaytseva宰采瓦 feud.Barony = &宰采瓦ZaytsevaBarony{}

func init() {
	f := BZaytseva宰采瓦.(*宰采瓦ZaytsevaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaytseva",
		TitleName: "宰采瓦",
		TitleCode: "b_zaytseva",
	}
}
