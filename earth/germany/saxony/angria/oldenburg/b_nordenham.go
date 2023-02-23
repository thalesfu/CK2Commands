package oldenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺登哈姆NordenhamBarony struct {
	feud.BaseBarony
}

var BNordenham诺登哈姆 feud.Barony = &诺登哈姆NordenhamBarony{}

func init() {
	f := BNordenham诺登哈姆.(*诺登哈姆NordenhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nordenham",
		TitleName: "诺登哈姆",
		TitleCode: "b_nordenham",
	}
}
