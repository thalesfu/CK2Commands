package dodugu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂希利特TichilitBarony struct {
	feud.BaseBarony
}

var BTichilit蒂希利特 feud.Barony = &蒂希利特TichilitBarony{}

func init() {
	f := BTichilit蒂希利特.(*蒂希利特TichilitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tichilit",
		TitleName: "蒂希利特",
		TitleCode: "b_tichilit",
	}
}
