package aqtobe

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔塔BurtaBarony struct {
	feud.BaseBarony
}

var BBurta布尔塔 feud.Barony = &布尔塔BurtaBarony{}

func init() {
    f := BBurta布尔塔.(*布尔塔BurtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burta",
		TitleName: "布尔塔",
		TitleCode: "b_burta",
	}
}
