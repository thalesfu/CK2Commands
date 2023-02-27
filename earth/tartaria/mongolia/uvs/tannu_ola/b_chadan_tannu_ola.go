package tannu_ola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰丹Chadan_tannu_olaBarony struct {
	feud.BaseBarony
}

var BChadan_tannu_ola恰丹 feud.Barony = &恰丹Chadan_tannu_olaBarony{}

func init() {
    f := BChadan_tannu_ola恰丹.(*恰丹Chadan_tannu_olaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chadan_tannu_ola",
		TitleName: "恰丹",
		TitleCode: "b_chadan_tannu_ola",
	}
}
