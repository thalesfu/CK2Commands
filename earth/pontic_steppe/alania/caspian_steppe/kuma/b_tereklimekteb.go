package kuma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷列克利_梅克捷布TereklimektebBarony struct {
	feud.BaseBarony
}

var BTereklimekteb捷列克利_梅克捷布 feud.Barony = &捷列克利_梅克捷布TereklimektebBarony{}

func init() {
    f := BTereklimekteb捷列克利_梅克捷布.(*捷列克利_梅克捷布TereklimektebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tereklimekteb",
		TitleName: "捷列克利_梅克捷布",
		TitleCode: "b_tereklimekteb",
	}
}
