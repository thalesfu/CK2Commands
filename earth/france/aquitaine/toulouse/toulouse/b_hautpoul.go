package toulouse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧特普尔HautpoulBarony struct {
	feud.BaseBarony
}

var BHautpoul欧特普尔 feud.Barony = &欧特普尔HautpoulBarony{}

func init() {
	f := BHautpoul欧特普尔.(*欧特普尔HautpoulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hautpoul",
		TitleName: "欧特普尔",
		TitleCode: "b_hautpoul",
	}
}
