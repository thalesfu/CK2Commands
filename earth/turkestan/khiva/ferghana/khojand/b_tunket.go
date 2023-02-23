package khojand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 腾凯特TunketBarony struct {
	feud.BaseBarony
}

var BTunket腾凯特 feud.Barony = &腾凯特TunketBarony{}

func init() {
	f := BTunket腾凯特.(*腾凯特TunketBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tunket",
		TitleName: "腾凯特",
		TitleCode: "b_tunket",
	}
}
