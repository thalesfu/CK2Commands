package valabhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古蒂迦GundigarBarony struct {
	feud.BaseBarony
}

var BGundigar古蒂迦 feud.Barony = &古蒂迦GundigarBarony{}

func init() {
	f := BGundigar古蒂迦.(*古蒂迦GundigarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gundigar",
		TitleName: "古蒂迦",
		TitleCode: "b_gundigar",
	}
}
