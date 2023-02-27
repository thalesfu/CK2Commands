package syrt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔希亚TashiaBarony struct {
	feud.BaseBarony
}

var BTashia塔希亚 feud.Barony = &塔希亚TashiaBarony{}

func init() {
    f := BTashia塔希亚.(*塔希亚TashiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tashia",
		TitleName: "塔希亚",
		TitleCode: "b_tashia",
	}
}
