package infa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉巴特RabatBarony struct {
	feud.BaseBarony
}

var BRabat拉巴特 feud.Barony = &拉巴特RabatBarony{}

func init() {
	f := BRabat拉巴特.(*拉巴特RabatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rabat",
		TitleName: "拉巴特",
		TitleCode: "b_rabat",
	}
}
