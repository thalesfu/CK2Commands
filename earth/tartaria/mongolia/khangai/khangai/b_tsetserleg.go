package khangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 车车尔勒格TsetserlegBarony struct {
	feud.BaseBarony
}

var BTsetserleg车车尔勒格 feud.Barony = &车车尔勒格TsetserlegBarony{}

func init() {
	f := BTsetserleg车车尔勒格.(*车车尔勒格TsetserlegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsetserleg",
		TitleName: "车车尔勒格",
		TitleCode: "b_tsetserleg",
	}
}
