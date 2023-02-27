package varadzin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉皮纳KrapinaBarony struct {
	feud.BaseBarony
}

var BKrapina克拉皮纳 feud.Barony = &克拉皮纳KrapinaBarony{}

func init() {
    f := BKrapina克拉皮纳.(*克拉皮纳KrapinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krapina",
		TitleName: "克拉皮纳",
		TitleCode: "b_krapina",
	}
}
