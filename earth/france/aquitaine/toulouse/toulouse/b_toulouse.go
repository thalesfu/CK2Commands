package toulouse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图卢兹ToulouseBarony struct {
	feud.BaseBarony
}

var BToulouse图卢兹 feud.Barony = &图卢兹ToulouseBarony{}

func init() {
	f := BToulouse图卢兹.(*图卢兹ToulouseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toulouse",
		TitleName: "图卢兹",
		TitleCode: "b_toulouse",
	}
}
