package la_mancha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉索纳TarazonaBarony struct {
	feud.BaseBarony
}

var BTarazona塔拉索纳 feud.Barony = &塔拉索纳TarazonaBarony{}

func init() {
    f := BTarazona塔拉索纳.(*塔拉索纳TarazonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarazona",
		TitleName: "塔拉索纳",
		TitleCode: "b_tarazona",
	}
}
