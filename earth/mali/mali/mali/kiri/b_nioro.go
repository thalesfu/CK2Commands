package kiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼奥罗NioroBarony struct {
	feud.BaseBarony
}

var BNioro尼奥罗 feud.Barony = &尼奥罗NioroBarony{}

func init() {
    f := BNioro尼奥罗.(*尼奥罗NioroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nioro",
		TitleName: "尼奥罗",
		TitleCode: "b_nioro",
	}
}
