package lancaster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索利SawleyBarony struct {
	feud.BaseBarony
}

var BSawley索利 feud.Barony = &索利SawleyBarony{}

func init() {
	f := BSawley索利.(*索利SawleyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sawley",
		TitleName: "索利",
		TitleCode: "b_sawley",
	}
}
