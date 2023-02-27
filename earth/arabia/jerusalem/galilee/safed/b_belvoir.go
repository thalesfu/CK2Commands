package safed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔沃BelvoirBarony struct {
	feud.BaseBarony
}

var BBelvoir贝尔沃 feud.Barony = &贝尔沃BelvoirBarony{}

func init() {
    f := BBelvoir贝尔沃.(*贝尔沃BelvoirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belvoir",
		TitleName: "贝尔沃",
		TitleCode: "b_belvoir",
	}
}
