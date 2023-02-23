package vastergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卡伦达SkalundaBarony struct {
	feud.BaseBarony
}

var BSkalunda斯卡伦达 feud.Barony = &斯卡伦达SkalundaBarony{}

func init() {
	f := BSkalunda斯卡伦达.(*斯卡伦达SkalundaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skalunda",
		TitleName: "斯卡伦达",
		TitleCode: "b_skalunda",
	}
}
