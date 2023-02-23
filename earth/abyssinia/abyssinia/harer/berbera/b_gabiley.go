package berbera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖比莱GabileyBarony struct {
	feud.BaseBarony
}

var BGabiley盖比莱 feud.Barony = &盖比莱GabileyBarony{}

func init() {
	f := BGabiley盖比莱.(*盖比莱GabileyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gabiley",
		TitleName: "盖比莱",
		TitleCode: "b_gabiley",
	}
}
