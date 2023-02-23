package ostfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃森斯EsensBarony struct {
	feud.BaseBarony
}

var BEsens埃森斯 feud.Barony = &埃森斯EsensBarony{}

func init() {
	f := BEsens埃森斯.(*埃森斯EsensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esens",
		TitleName: "埃森斯",
		TitleCode: "b_esens",
	}
}
