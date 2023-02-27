package vikramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗讫罗摩补罗VikramapuraBarony struct {
	feud.BaseBarony
}

var BVikramapura毗讫罗摩补罗 feud.Barony = &毗讫罗摩补罗VikramapuraBarony{}

func init() {
    f := BVikramapura毗讫罗摩补罗.(*毗讫罗摩补罗VikramapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vikramapura",
		TitleName: "毗讫罗摩补罗",
		TitleCode: "b_vikramapura",
	}
}
