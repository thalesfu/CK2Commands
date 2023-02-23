package bayuda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅拉格AlmeraghBarony struct {
	feud.BaseBarony
}

var BAlmeragh梅拉格 feud.Barony = &梅拉格AlmeraghBarony{}

func init() {
	f := BAlmeragh梅拉格.(*梅拉格AlmeraghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almeragh",
		TitleName: "梅拉格",
		TitleCode: "b_almeragh",
	}
}
