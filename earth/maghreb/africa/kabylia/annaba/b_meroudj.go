package annaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅罗贾MeroudjBarony struct {
	feud.BaseBarony
}

var BMeroudj梅罗贾 feud.Barony = &梅罗贾MeroudjBarony{}

func init() {
	f := BMeroudj梅罗贾.(*梅罗贾MeroudjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meroudj",
		TitleName: "梅罗贾",
		TitleCode: "b_meroudj",
	}
}
