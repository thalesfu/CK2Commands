package kara_kum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科切夫卡KochevkaBarony struct {
	feud.BaseBarony
}

var BKochevka科切夫卡 feud.Barony = &科切夫卡KochevkaBarony{}

func init() {
    f := BKochevka科切夫卡.(*科切夫卡KochevkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kochevka",
		TitleName: "科切夫卡",
		TitleCode: "b_kochevka",
	}
}
