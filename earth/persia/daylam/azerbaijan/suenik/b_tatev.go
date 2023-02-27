package suenik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔特夫TatevBarony struct {
	feud.BaseBarony
}

var BTatev塔特夫 feud.Barony = &塔特夫TatevBarony{}

func init() {
    f := BTatev塔特夫.(*塔特夫TatevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tatev",
		TitleName: "塔特夫",
		TitleCode: "b_tatev",
	}
}
