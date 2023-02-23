package tortosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈瓦比AlkhawabiBarony struct {
	feud.BaseBarony
}

var BAlkhawabi哈瓦比 feud.Barony = &哈瓦比AlkhawabiBarony{}

func init() {
	f := BAlkhawabi哈瓦比.(*哈瓦比AlkhawabiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alkhawabi",
		TitleName: "哈瓦比",
		TitleCode: "b_alkhawabi",
	}
}
