package cuenca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌克莱斯UclesBarony struct {
	feud.BaseBarony
}

var BUcles乌克莱斯 feud.Barony = &乌克莱斯UclesBarony{}

func init() {
	f := BUcles乌克莱斯.(*乌克莱斯UclesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ucles",
		TitleName: "乌克莱斯",
		TitleCode: "b_ucles",
	}
}
