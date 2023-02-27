package rafha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朱迈马AljumaymahBarony struct {
	feud.BaseBarony
}

var BAljumaymah朱迈马 feud.Barony = &朱迈马AljumaymahBarony{}

func init() {
    f := BAljumaymah朱迈马.(*朱迈马AljumaymahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aljumaymah",
		TitleName: "朱迈马",
		TitleCode: "b_aljumaymah",
	}
}
