package nikopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝莱内BeleneBarony struct {
	feud.BaseBarony
}

var BBelene贝莱内 feud.Barony = &贝莱内BeleneBarony{}

func init() {
	f := BBelene贝莱内.(*贝莱内BeleneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belene",
		TitleName: "贝莱内",
		TitleCode: "b_belene",
	}
}
