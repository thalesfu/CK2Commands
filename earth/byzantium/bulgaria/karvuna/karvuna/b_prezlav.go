package karvuna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普雷斯拉夫PrezlavBarony struct {
	feud.BaseBarony
}

var BPrezlav普雷斯拉夫 feud.Barony = &普雷斯拉夫PrezlavBarony{}

func init() {
    f := BPrezlav普雷斯拉夫.(*普雷斯拉夫PrezlavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prezlav",
		TitleName: "普雷斯拉夫",
		TitleCode: "b_prezlav",
	}
}
