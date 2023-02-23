package devon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达特茅斯DartmouthBarony struct {
	feud.BaseBarony
}

var BDartmouth达特茅斯 feud.Barony = &达特茅斯DartmouthBarony{}

func init() {
	f := BDartmouth达特茅斯.(*达特茅斯DartmouthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dartmouth",
		TitleName: "达特茅斯",
		TitleCode: "b_dartmouth",
	}
}
