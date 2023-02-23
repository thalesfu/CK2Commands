package karvuna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔武纳KarvunaBarony struct {
	feud.BaseBarony
}

var BKarvuna卡尔武纳 feud.Barony = &卡尔武纳KarvunaBarony{}

func init() {
	f := BKarvuna卡尔武纳.(*卡尔武纳KarvunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karvuna",
		TitleName: "卡尔武纳",
		TitleCode: "b_karvuna",
	}
}
