package saray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰尔加DzhelgaBarony struct {
	feud.BaseBarony
}

var BDzhelga杰尔加 feud.Barony = &杰尔加DzhelgaBarony{}

func init() {
    f := BDzhelga杰尔加.(*杰尔加DzhelgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dzhelga",
		TitleName: "杰尔加",
		TitleCode: "b_dzhelga",
	}
}
