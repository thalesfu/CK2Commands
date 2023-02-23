package ennedi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯布卡比亚KebkabiyaBarony struct {
	feud.BaseBarony
}

var BKebkabiya凯布卡比亚 feud.Barony = &凯布卡比亚KebkabiyaBarony{}

func init() {
	f := BKebkabiya凯布卡比亚.(*凯布卡比亚KebkabiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kebkabiya",
		TitleName: "凯布卡比亚",
		TitleCode: "b_kebkabiya",
	}
}
