package kirkuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基夫里KifriBarony struct {
	feud.BaseBarony
}

var BKifri基夫里 feud.Barony = &基夫里KifriBarony{}

func init() {
	f := BKifri基夫里.(*基夫里KifriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kifri",
		TitleName: "基夫里",
		TitleCode: "b_kifri",
	}
}
