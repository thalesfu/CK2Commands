package podlasie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德罗希琴DrohiczynBarony struct {
	feud.BaseBarony
}

var BDrohiczyn德罗希琴 feud.Barony = &德罗希琴DrohiczynBarony{}

func init() {
    f := BDrohiczyn德罗希琴.(*德罗希琴DrohiczynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drohiczyn",
		TitleName: "德罗希琴",
		TitleCode: "b_drohiczyn",
	}
}
