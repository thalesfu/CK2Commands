package draa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坦尚南特TamchennenntBarony struct {
	feud.BaseBarony
}

var BTamchennennt坦尚南特 feud.Barony = &坦尚南特TamchennenntBarony{}

func init() {
	f := BTamchennennt坦尚南特.(*坦尚南特TamchennenntBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamchennennt",
		TitleName: "坦尚南特",
		TitleCode: "b_tamchennennt",
	}
}
