package tarvagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷列霍尔Tere_kholBarony struct {
	feud.BaseBarony
}

var BTere_khol捷列霍尔 feud.Barony = &捷列霍尔Tere_kholBarony{}

func init() {
    f := BTere_khol捷列霍尔.(*捷列霍尔Tere_kholBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tere_khol",
		TitleName: "捷列霍尔",
		TitleCode: "b_tere_khol",
	}
}
