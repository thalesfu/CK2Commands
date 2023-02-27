package mainz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 美因茨MainzBarony struct {
	feud.BaseBarony
}

var BMainz美因茨 feud.Barony = &美因茨MainzBarony{}

func init() {
    f := BMainz美因茨.(*美因茨MainzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mainz",
		TitleName: "美因茨",
		TitleCode: "b_mainz",
	}
}
