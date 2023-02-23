package abauj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔尼亚TurnaBarony struct {
	feud.BaseBarony
}

var BTurna图尔尼亚 feud.Barony = &图尔尼亚TurnaBarony{}

func init() {
	f := BTurna图尔尼亚.(*图尔尼亚TurnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turna",
		TitleName: "图尔尼亚",
		TitleCode: "b_turna",
	}
}
