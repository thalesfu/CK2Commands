package savoie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克吕斯ClusesBarony struct {
	feud.BaseBarony
}

var BCluses克吕斯 feud.Barony = &克吕斯ClusesBarony{}

func init() {
	f := BCluses克吕斯.(*克吕斯ClusesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cluses",
		TitleName: "克吕斯",
		TitleCode: "b_cluses",
	}
}
