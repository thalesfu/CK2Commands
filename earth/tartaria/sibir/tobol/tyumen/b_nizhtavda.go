package tyumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下塔夫达NizhtavdaBarony struct {
	feud.BaseBarony
}

var BNizhtavda下塔夫达 feud.Barony = &下塔夫达NizhtavdaBarony{}

func init() {
	f := BNizhtavda下塔夫达.(*下塔夫达NizhtavdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizhtavda",
		TitleName: "下塔夫达",
		TitleCode: "b_nizhtavda",
	}
}
