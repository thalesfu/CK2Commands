package madurai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提卢波兰贡愣TiruparankunramBarony struct {
	feud.BaseBarony
}

var BTiruparankunram提卢波兰贡愣 feud.Barony = &提卢波兰贡愣TiruparankunramBarony{}

func init() {
    f := BTiruparankunram提卢波兰贡愣.(*提卢波兰贡愣TiruparankunramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiruparankunram",
		TitleName: "提卢波兰贡愣",
		TitleCode: "b_tiruparankunram",
	}
}
