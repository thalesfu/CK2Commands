package podlasie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉伊格鲁德RajgrodBarony struct {
	feud.BaseBarony
}

var BRajgrod拉伊格鲁德 feud.Barony = &拉伊格鲁德RajgrodBarony{}

func init() {
    f := BRajgrod拉伊格鲁德.(*拉伊格鲁德RajgrodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajgrod",
		TitleName: "拉伊格鲁德",
		TitleCode: "b_rajgrod",
	}
}
