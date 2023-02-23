package strathearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图里巴丁TullibardineBarony struct {
	feud.BaseBarony
}

var BTullibardine图里巴丁 feud.Barony = &图里巴丁TullibardineBarony{}

func init() {
	f := BTullibardine图里巴丁.(*图里巴丁TullibardineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tullibardine",
		TitleName: "图里巴丁",
		TitleCode: "b_tullibardine",
	}
}
