package kolzum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陶菲克TaufiqBarony struct {
	feud.BaseBarony
}

var BTaufiq陶菲克 feud.Barony = &陶菲克TaufiqBarony{}

func init() {
	f := BTaufiq陶菲克.(*陶菲克TaufiqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taufiq",
		TitleName: "陶菲克",
		TitleCode: "b_taufiq",
	}
}
