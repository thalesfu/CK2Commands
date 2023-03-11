package sieradzko_leczyckie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈斯蒂宁GostyninBarony struct {
	feud.BaseBarony
}

var BGostynin戈斯蒂宁 feud.Barony = &戈斯蒂宁GostyninBarony{}

func init() {
    f := BGostynin戈斯蒂宁.(*戈斯蒂宁GostyninBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gostynin",
		TitleName: "戈斯蒂宁",
		TitleCode: "b_gostynin",
	}
}
