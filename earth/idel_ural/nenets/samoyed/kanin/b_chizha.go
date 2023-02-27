package kanin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇扎ChizhaBarony struct {
	feud.BaseBarony
}

var BChizha奇扎 feud.Barony = &奇扎ChizhaBarony{}

func init() {
    f := BChizha奇扎.(*奇扎ChizhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chizha",
		TitleName: "奇扎",
		TitleCode: "b_chizha",
	}
}
