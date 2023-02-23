package foggia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博维诺BovinoBarony struct {
	feud.BaseBarony
}

var BBovino博维诺 feud.Barony = &博维诺BovinoBarony{}

func init() {
	f := BBovino博维诺.(*博维诺BovinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bovino",
		TitleName: "博维诺",
		TitleCode: "b_bovino",
	}
}
