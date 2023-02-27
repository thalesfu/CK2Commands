package gizeh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨卡拉SaqqaraBarony struct {
	feud.BaseBarony
}

var BSaqqara萨卡拉 feud.Barony = &萨卡拉SaqqaraBarony{}

func init() {
    f := BSaqqara萨卡拉.(*萨卡拉SaqqaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saqqara",
		TitleName: "萨卡拉",
		TitleCode: "b_saqqara",
	}
}
