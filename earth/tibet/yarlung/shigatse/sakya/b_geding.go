package sakya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉定GedingBarony struct {
	feud.BaseBarony
}

var BGeding吉定 feud.Barony = &吉定GedingBarony{}

func init() {
    f := BGeding吉定.(*吉定GedingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geding",
		TitleName: "吉定",
		TitleCode: "b_geding",
	}
}
