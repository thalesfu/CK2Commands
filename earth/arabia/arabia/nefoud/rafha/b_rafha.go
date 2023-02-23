package rafha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉夫哈RafhaBarony struct {
	feud.BaseBarony
}

var BRafha拉夫哈 feud.Barony = &拉夫哈RafhaBarony{}

func init() {
	f := BRafha拉夫哈.(*拉夫哈RafhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rafha",
		TitleName: "拉夫哈",
		TitleCode: "b_rafha",
	}
}
