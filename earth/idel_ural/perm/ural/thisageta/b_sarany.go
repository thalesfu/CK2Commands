package thisageta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉内SaranyBarony struct {
	feud.BaseBarony
}

var BSarany萨拉内 feud.Barony = &萨拉内SaranyBarony{}

func init() {
    f := BSarany萨拉内.(*萨拉内SaranyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarany",
		TitleName: "萨拉内",
		TitleCode: "b_sarany",
	}
}
