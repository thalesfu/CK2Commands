package reims

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃佩尔奈EpernayBarony struct {
	feud.BaseBarony
}

var BEpernay埃佩尔奈 feud.Barony = &埃佩尔奈EpernayBarony{}

func init() {
    f := BEpernay埃佩尔奈.(*埃佩尔奈EpernayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "epernay",
		TitleName: "埃佩尔奈",
		TitleCode: "b_epernay",
	}
}
