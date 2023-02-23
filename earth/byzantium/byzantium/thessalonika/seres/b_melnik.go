package seres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔尼克MelnikBarony struct {
	feud.BaseBarony
}

var BMelnik梅尔尼克 feud.Barony = &梅尔尼克MelnikBarony{}

func init() {
	f := BMelnik梅尔尼克.(*梅尔尼克MelnikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melnik",
		TitleName: "梅尔尼克",
		TitleCode: "b_melnik",
	}
}
