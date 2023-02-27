package khopyor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坦波夫TambovBarony struct {
	feud.BaseBarony
}

var BTambov坦波夫 feud.Barony = &坦波夫TambovBarony{}

func init() {
    f := BTambov坦波夫.(*坦波夫TambovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tambov",
		TitleName: "坦波夫",
		TitleCode: "b_tambov",
	}
}
