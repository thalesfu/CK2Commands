package khiva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔赫塔TahtaBarony struct {
	feud.BaseBarony
}

var BTahta塔赫塔 feud.Barony = &塔赫塔TahtaBarony{}

func init() {
    f := BTahta塔赫塔.(*塔赫塔TahtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tahta",
		TitleName: "塔赫塔",
		TitleCode: "b_tahta",
	}
}
