package zanjan_abhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴AbaBarony struct {
	feud.BaseBarony
}

var BAba阿巴 feud.Barony = &阿巴AbaBarony{}

func init() {
    f := BAba阿巴.(*阿巴AbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aba",
		TitleName: "阿巴",
		TitleCode: "b_aba",
	}
}
