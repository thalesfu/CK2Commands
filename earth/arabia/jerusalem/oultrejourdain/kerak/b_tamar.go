package kerak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔马林TamarBarony struct {
	feud.BaseBarony
}

var BTamar塔马林 feud.Barony = &塔马林TamarBarony{}

func init() {
    f := BTamar塔马林.(*塔马林TamarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamar",
		TitleName: "塔马林",
		TitleCode: "b_tamar",
	}
}
