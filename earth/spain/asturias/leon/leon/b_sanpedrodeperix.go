package leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣佩德罗德佩里克斯SanpedrodeperixBarony struct {
	feud.BaseBarony
}

var BSanpedrodeperix圣佩德罗德佩里克斯 feud.Barony = &圣佩德罗德佩里克斯SanpedrodeperixBarony{}

func init() {
	f := BSanpedrodeperix圣佩德罗德佩里克斯.(*圣佩德罗德佩里克斯SanpedrodeperixBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanpedrodeperix",
		TitleName: "圣佩德罗德佩里克斯",
		TitleCode: "b_sanpedrodeperix",
	}
}
