package srirangapatna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝拉普拉BelapuraBarony struct {
	feud.BaseBarony
}

var BBelapura贝拉普拉 feud.Barony = &贝拉普拉BelapuraBarony{}

func init() {
    f := BBelapura贝拉普拉.(*贝拉普拉BelapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belapura",
		TitleName: "贝拉普拉",
		TitleCode: "b_belapura",
	}
}
