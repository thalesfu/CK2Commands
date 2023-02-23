package vadodara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达婆伐底DarbhavatiBarony struct {
	feud.BaseBarony
}

var BDarbhavati达婆伐底 feud.Barony = &达婆伐底DarbhavatiBarony{}

func init() {
	f := BDarbhavati达婆伐底.(*达婆伐底DarbhavatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darbhavati",
		TitleName: "达婆伐底",
		TitleCode: "b_darbhavati",
	}
}
