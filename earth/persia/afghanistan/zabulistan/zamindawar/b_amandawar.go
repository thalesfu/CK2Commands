package zamindawar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿曼达瓦尔AmandawarBarony struct {
	feud.BaseBarony
}

var BAmandawar阿曼达瓦尔 feud.Barony = &阿曼达瓦尔AmandawarBarony{}

func init() {
    f := BAmandawar阿曼达瓦尔.(*阿曼达瓦尔AmandawarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amandawar",
		TitleName: "阿曼达瓦尔",
		TitleCode: "b_amandawar",
	}
}
