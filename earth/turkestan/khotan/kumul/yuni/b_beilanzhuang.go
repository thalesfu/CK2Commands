package yuni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 北兰庄BeilanzhuangBarony struct {
	feud.BaseBarony
}

var BBeilanzhuang北兰庄 feud.Barony = &北兰庄BeilanzhuangBarony{}

func init() {
    f := BBeilanzhuang北兰庄.(*北兰庄BeilanzhuangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beilanzhuang",
		TitleName: "北兰庄",
		TitleCode: "b_beilanzhuang",
	}
}
