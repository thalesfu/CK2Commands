package epeiros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉罗卡斯特GjirokasterBarony struct {
	feud.BaseBarony
}

var BGjirokaster吉罗卡斯特 feud.Barony = &吉罗卡斯特GjirokasterBarony{}

func init() {
	f := BGjirokaster吉罗卡斯特.(*吉罗卡斯特GjirokasterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gjirokaster",
		TitleName: "吉罗卡斯特",
		TitleCode: "b_gjirokaster",
	}
}
