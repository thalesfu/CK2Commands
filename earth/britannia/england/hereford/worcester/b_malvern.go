package worcester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔文MalvernBarony struct {
	feud.BaseBarony
}

var BMalvern莫尔文 feud.Barony = &莫尔文MalvernBarony{}

func init() {
	f := BMalvern莫尔文.(*莫尔文MalvernBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malvern",
		TitleName: "莫尔文",
		TitleCode: "b_malvern",
	}
}
