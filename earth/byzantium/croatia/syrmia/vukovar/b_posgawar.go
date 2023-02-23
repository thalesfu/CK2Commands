package vukovar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波斯加瓦尔PosgawarBarony struct {
	feud.BaseBarony
}

var BPosgawar波斯加瓦尔 feud.Barony = &波斯加瓦尔PosgawarBarony{}

func init() {
	f := BPosgawar波斯加瓦尔.(*波斯加瓦尔PosgawarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "posgawar",
		TitleName: "波斯加瓦尔",
		TitleCode: "b_posgawar",
	}
}
