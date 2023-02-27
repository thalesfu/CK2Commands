package navarra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 潘普洛纳PamplonaBarony struct {
	feud.BaseBarony
}

var BPamplona潘普洛纳 feud.Barony = &潘普洛纳PamplonaBarony{}

func init() {
    f := BPamplona潘普洛纳.(*潘普洛纳PamplonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pamplona",
		TitleName: "潘普洛纳",
		TitleCode: "b_pamplona",
	}
}
