package soli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯雷布雷尼克SrebrenikBarony struct {
	feud.BaseBarony
}

var BSrebrenik斯雷布雷尼克 feud.Barony = &斯雷布雷尼克SrebrenikBarony{}

func init() {
    f := BSrebrenik斯雷布雷尼克.(*斯雷布雷尼克SrebrenikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srebrenik",
		TitleName: "斯雷布雷尼克",
		TitleCode: "b_srebrenik",
	}
}
