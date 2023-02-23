package suakin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德尔金DerkinBarony struct {
	feud.BaseBarony
}

var BDerkin德尔金 feud.Barony = &德尔金DerkinBarony{}

func init() {
	f := BDerkin德尔金.(*德尔金DerkinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derkin",
		TitleName: "德尔金",
		TitleCode: "b_derkin",
	}
}
