package gevaudan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔沃若勒MarvejolsBarony struct {
	feud.BaseBarony
}

var BMarvejols马尔沃若勒 feud.Barony = &马尔沃若勒MarvejolsBarony{}

func init() {
    f := BMarvejols马尔沃若勒.(*马尔沃若勒MarvejolsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marvejols",
		TitleName: "马尔沃若勒",
		TitleCode: "b_marvejols",
	}
}
