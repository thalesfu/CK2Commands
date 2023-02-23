package seletyteniz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢列特捷尼兹SeletytenizBarony struct {
	feud.BaseBarony
}

var BSeletyteniz谢列特捷尼兹 feud.Barony = &谢列特捷尼兹SeletytenizBarony{}

func init() {
	f := BSeletyteniz谢列特捷尼兹.(*谢列特捷尼兹SeletytenizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seletyteniz",
		TitleName: "谢列特捷尼兹",
		TitleCode: "b_seletyteniz",
	}
}
