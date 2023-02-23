package parma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉斯佩齐亚LaspeziaBarony struct {
	feud.BaseBarony
}

var BLaspezia拉斯佩齐亚 feud.Barony = &拉斯佩齐亚LaspeziaBarony{}

func init() {
	f := BLaspezia拉斯佩齐亚.(*拉斯佩齐亚LaspeziaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laspezia",
		TitleName: "拉斯佩齐亚",
		TitleCode: "b_laspezia",
	}
}
