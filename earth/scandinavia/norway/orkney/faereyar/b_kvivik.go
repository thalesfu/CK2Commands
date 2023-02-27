package faereyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克维维克KvivikBarony struct {
	feud.BaseBarony
}

var BKvivik克维维克 feud.Barony = &克维维克KvivikBarony{}

func init() {
    f := BKvivik克维维克.(*克维维克KvivikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kvivik",
		TitleName: "克维维克",
		TitleCode: "b_kvivik",
	}
}
