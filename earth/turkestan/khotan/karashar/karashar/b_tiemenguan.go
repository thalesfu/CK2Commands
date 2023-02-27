package karashar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 铁门关TiemenguanBarony struct {
	feud.BaseBarony
}

var BTiemenguan铁门关 feud.Barony = &铁门关TiemenguanBarony{}

func init() {
    f := BTiemenguan铁门关.(*铁门关TiemenguanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiemenguan",
		TitleName: "铁门关",
		TitleCode: "b_tiemenguan",
	}
}
