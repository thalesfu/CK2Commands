package vermandois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈希GauchyBarony struct {
	feud.BaseBarony
}

var BGauchy戈希 feud.Barony = &戈希GauchyBarony{}

func init() {
	f := BGauchy戈希.(*戈希GauchyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gauchy",
		TitleName: "戈希",
		TitleCode: "b_gauchy",
	}
}
