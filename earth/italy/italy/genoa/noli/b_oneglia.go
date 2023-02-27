package noli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥内利亚OnegliaBarony struct {
	feud.BaseBarony
}

var BOneglia奥内利亚 feud.Barony = &奥内利亚OnegliaBarony{}

func init() {
    f := BOneglia奥内利亚.(*奥内利亚OnegliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oneglia",
		TitleName: "奥内利亚",
		TitleCode: "b_oneglia",
	}
}
