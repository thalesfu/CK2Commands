package lyon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昂斯AnseBarony struct {
	feud.BaseBarony
}

var BAnse昂斯 feud.Barony = &昂斯AnseBarony{}

func init() {
	f := BAnse昂斯.(*昂斯AnseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anse",
		TitleName: "昂斯",
		TitleCode: "b_anse",
	}
}
