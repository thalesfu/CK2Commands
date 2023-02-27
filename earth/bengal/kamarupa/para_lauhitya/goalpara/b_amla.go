package goalpara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆拉AmlaBarony struct {
	feud.BaseBarony
}

var BAmla阿姆拉 feud.Barony = &阿姆拉AmlaBarony{}

func init() {
    f := BAmla阿姆拉.(*阿姆拉AmlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amla",
		TitleName: "阿姆拉",
		TitleCode: "b_amla",
	}
}
