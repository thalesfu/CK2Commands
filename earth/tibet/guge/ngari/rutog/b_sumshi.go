package rutog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松西SumshiBarony struct {
	feud.BaseBarony
}

var BSumshi松西 feud.Barony = &松西SumshiBarony{}

func init() {
    f := BSumshi松西.(*松西SumshiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sumshi",
		TitleName: "松西",
		TitleCode: "b_sumshi",
	}
}
