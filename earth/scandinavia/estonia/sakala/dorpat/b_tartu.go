package dorpat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔图TartuBarony struct {
	feud.BaseBarony
}

var BTartu塔尔图 feud.Barony = &塔尔图TartuBarony{}

func init() {
	f := BTartu塔尔图.(*塔尔图TartuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tartu",
		TitleName: "塔尔图",
		TitleCode: "b_tartu",
	}
}
