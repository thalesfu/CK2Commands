package urzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿亚古兹AyagozBarony struct {
	feud.BaseBarony
}

var BAyagoz阿亚古兹 feud.Barony = &阿亚古兹AyagozBarony{}

func init() {
	f := BAyagoz阿亚古兹.(*阿亚古兹AyagozBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ayagoz",
		TitleName: "阿亚古兹",
		TitleCode: "b_ayagoz",
	}
}
