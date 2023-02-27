package feher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维佐克瑙VizaknaBarony struct {
	feud.BaseBarony
}

var BVizakna维佐克瑙 feud.Barony = &维佐克瑙VizaknaBarony{}

func init() {
    f := BVizakna维佐克瑙.(*维佐克瑙VizaknaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vizakna",
		TitleName: "维佐克瑙",
		TitleCode: "b_vizakna",
	}
}
