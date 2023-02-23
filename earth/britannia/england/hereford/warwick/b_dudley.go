package warwick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达德利DudleyBarony struct {
	feud.BaseBarony
}

var BDudley达德利 feud.Barony = &达德利DudleyBarony{}

func init() {
	f := BDudley达德利.(*达德利DudleyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dudley",
		TitleName: "达德利",
		TitleCode: "b_dudley",
	}
}
