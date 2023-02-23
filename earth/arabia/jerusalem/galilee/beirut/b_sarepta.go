package beirut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 撒勒法SareptaBarony struct {
	feud.BaseBarony
}

var BSarepta撒勒法 feud.Barony = &撒勒法SareptaBarony{}

func init() {
	f := BSarepta撒勒法.(*撒勒法SareptaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarepta",
		TitleName: "撒勒法",
		TitleCode: "b_sarepta",
	}
}
