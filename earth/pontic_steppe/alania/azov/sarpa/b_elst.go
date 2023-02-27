package sarpa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶尔斯特ElstBarony struct {
	feud.BaseBarony
}

var BElst叶尔斯特 feud.Barony = &叶尔斯特ElstBarony{}

func init() {
    f := BElst叶尔斯特.(*叶尔斯特ElstBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elst",
		TitleName: "叶尔斯特",
		TitleCode: "b_elst",
	}
}
