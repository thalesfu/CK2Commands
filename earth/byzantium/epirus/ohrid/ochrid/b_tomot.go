package ochrid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 土默特TomotBarony struct {
	feud.BaseBarony
}

var BTomot土默特 feud.Barony = &土默特TomotBarony{}

func init() {
	f := BTomot土默特.(*土默特TomotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tomot",
		TitleName: "土默特",
		TitleCode: "b_tomot",
	}
}
