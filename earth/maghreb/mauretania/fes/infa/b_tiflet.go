package infa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提夫莱特TifletBarony struct {
	feud.BaseBarony
}

var BTiflet提夫莱特 feud.Barony = &提夫莱特TifletBarony{}

func init() {
    f := BTiflet提夫莱特.(*提夫莱特TifletBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiflet",
		TitleName: "提夫莱特",
		TitleCode: "b_tiflet",
	}
}
