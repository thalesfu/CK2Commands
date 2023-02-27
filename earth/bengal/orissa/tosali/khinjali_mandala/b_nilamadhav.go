package khinjali_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼罗摩陀婆NilamadhavBarony struct {
	feud.BaseBarony
}

var BNilamadhav尼罗摩陀婆 feud.Barony = &尼罗摩陀婆NilamadhavBarony{}

func init() {
    f := BNilamadhav尼罗摩陀婆.(*尼罗摩陀婆NilamadhavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nilamadhav",
		TitleName: "尼罗摩陀婆",
		TitleCode: "b_nilamadhav",
	}
}
