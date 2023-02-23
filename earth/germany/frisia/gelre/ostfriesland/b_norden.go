package ostfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺登NordenBarony struct {
	feud.BaseBarony
}

var BNorden诺登 feud.Barony = &诺登NordenBarony{}

func init() {
	f := BNorden诺登.(*诺登NordenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "norden",
		TitleName: "诺登",
		TitleCode: "b_norden",
	}
}
