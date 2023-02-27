package memel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考普KaupBarony struct {
	feud.BaseBarony
}

var BKaup考普 feud.Barony = &考普KaupBarony{}

func init() {
    f := BKaup考普.(*考普KaupBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaup",
		TitleName: "考普",
		TitleCode: "b_kaup",
	}
}
