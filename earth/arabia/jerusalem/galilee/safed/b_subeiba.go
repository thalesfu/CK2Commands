package safed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宁录堡SubeibaBarony struct {
	feud.BaseBarony
}

var BSubeiba宁录堡 feud.Barony = &宁录堡SubeibaBarony{}

func init() {
    f := BSubeiba宁录堡.(*宁录堡SubeibaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "subeiba",
		TitleName: "宁录堡",
		TitleCode: "b_subeiba",
	}
}
