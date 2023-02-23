package foggia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉普拉SalaplaBarony struct {
	feud.BaseBarony
}

var BSalapla萨拉普拉 feud.Barony = &萨拉普拉SalaplaBarony{}

func init() {
	f := BSalapla萨拉普拉.(*萨拉普拉SalaplaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salapla",
		TitleName: "萨拉普拉",
		TitleCode: "b_salapla",
	}
}
