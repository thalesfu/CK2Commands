package bereg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 派赖切尼PerecsenyBarony struct {
	feud.BaseBarony
}

var BPerecseny派赖切尼 feud.Barony = &派赖切尼PerecsenyBarony{}

func init() {
    f := BPerecseny派赖切尼.(*派赖切尼PerecsenyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perecseny",
		TitleName: "派赖切尼",
		TitleCode: "b_perecseny",
	}
}
