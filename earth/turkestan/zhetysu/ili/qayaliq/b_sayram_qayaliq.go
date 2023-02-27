package qayaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛里木Sayram_qayaliqBarony struct {
	feud.BaseBarony
}

var BSayram_qayaliq赛里木 feud.Barony = &赛里木Sayram_qayaliqBarony{}

func init() {
    f := BSayram_qayaliq赛里木.(*赛里木Sayram_qayaliqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sayram_qayaliq",
		TitleName: "赛里木",
		TitleCode: "b_sayram_qayaliq",
	}
}
