package tobruk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾卡拉达拜AlqardabahBarony struct {
	feud.BaseBarony
}

var BAlqardabah艾卡拉达拜 feud.Barony = &艾卡拉达拜AlqardabahBarony{}

func init() {
    f := BAlqardabah艾卡拉达拜.(*艾卡拉达拜AlqardabahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alqardabah",
		TitleName: "艾卡拉达拜",
		TitleCode: "b_alqardabah",
	}
}
