package tavda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔夫达TavdaBarony struct {
	feud.BaseBarony
}

var BTavda塔夫达 feud.Barony = &塔夫达TavdaBarony{}

func init() {
	f := BTavda塔夫达.(*塔夫达TavdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tavda",
		TitleName: "塔夫达",
		TitleCode: "b_tavda",
	}
}
