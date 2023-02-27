package archa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特耶兹TreizBarony struct {
	feud.BaseBarony
}

var BTreiz特耶兹 feud.Barony = &特耶兹TreizBarony{}

func init() {
    f := BTreiz特耶兹.(*特耶兹TreizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "treiz",
		TitleName: "特耶兹",
		TitleCode: "b_treiz",
	}
}
