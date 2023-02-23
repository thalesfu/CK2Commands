package tsaparang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 噶大克GartokBarony struct {
	feud.BaseBarony
}

var BGartok噶大克 feud.Barony = &噶大克GartokBarony{}

func init() {
	f := BGartok噶大克.(*噶大克GartokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gartok",
		TitleName: "噶大克",
		TitleCode: "b_gartok",
	}
}
