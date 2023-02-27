package claudiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤利奥波利斯IuliopolisBarony struct {
	feud.BaseBarony
}

var BIuliopolis尤利奥波利斯 feud.Barony = &尤利奥波利斯IuliopolisBarony{}

func init() {
    f := BIuliopolis尤利奥波利斯.(*尤利奥波利斯IuliopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iuliopolis",
		TitleName: "尤利奥波利斯",
		TitleCode: "b_iuliopolis",
	}
}
