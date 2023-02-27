package cumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡莱尔CarlisleBarony struct {
	feud.BaseBarony
}

var BCarlisle卡莱尔 feud.Barony = &卡莱尔CarlisleBarony{}

func init() {
    f := BCarlisle卡莱尔.(*卡莱尔CarlisleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carlisle",
		TitleName: "卡莱尔",
		TitleCode: "b_carlisle",
	}
}
