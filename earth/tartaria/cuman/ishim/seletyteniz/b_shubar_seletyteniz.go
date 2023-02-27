package seletyteniz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒巴尔Shubar_seletytenizBarony struct {
	feud.BaseBarony
}

var BShubar_seletyteniz舒巴尔 feud.Barony = &舒巴尔Shubar_seletytenizBarony{}

func init() {
    f := BShubar_seletyteniz舒巴尔.(*舒巴尔Shubar_seletytenizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shubar_seletyteniz",
		TitleName: "舒巴尔",
		TitleCode: "b_shubar_seletyteniz",
	}
}
