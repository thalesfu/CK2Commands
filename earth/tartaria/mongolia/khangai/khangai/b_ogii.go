package khangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 额吉OgiiBarony struct {
	feud.BaseBarony
}

var BOgii额吉 feud.Barony = &额吉OgiiBarony{}

func init() {
    f := BOgii额吉.(*额吉OgiiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ogii",
		TitleName: "额吉",
		TitleCode: "b_ogii",
	}
}
