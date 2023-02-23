package geneve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日内瓦GeneveBarony struct {
	feud.BaseBarony
}

var BGeneve日内瓦 feud.Barony = &日内瓦GeneveBarony{}

func init() {
	f := BGeneve日内瓦.(*日内瓦GeneveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geneve",
		TitleName: "日内瓦",
		TitleCode: "b_geneve",
	}
}
