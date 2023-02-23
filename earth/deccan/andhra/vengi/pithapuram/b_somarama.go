package pithapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏马拉玛SomaramaBarony struct {
	feud.BaseBarony
}

var BSomarama苏马拉玛 feud.Barony = &苏马拉玛SomaramaBarony{}

func init() {
	f := BSomarama苏马拉玛.(*苏马拉玛SomaramaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "somarama",
		TitleName: "苏马拉玛",
		TitleCode: "b_somarama",
	}
}
