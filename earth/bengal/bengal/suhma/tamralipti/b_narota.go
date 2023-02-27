package tamralipti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那卢多NarotaBarony struct {
	feud.BaseBarony
}

var BNarota那卢多 feud.Barony = &那卢多NarotaBarony{}

func init() {
    f := BNarota那卢多.(*那卢多NarotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narota",
		TitleName: "那卢多",
		TitleCode: "b_narota",
	}
}
