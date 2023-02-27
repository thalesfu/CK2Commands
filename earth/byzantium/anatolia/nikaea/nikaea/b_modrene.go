package nikaea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫德里尼ModreneBarony struct {
	feud.BaseBarony
}

var BModrene莫德里尼 feud.Barony = &莫德里尼ModreneBarony{}

func init() {
    f := BModrene莫德里尼.(*莫德里尼ModreneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "modrene",
		TitleName: "莫德里尼",
		TitleCode: "b_modrene",
	}
}
