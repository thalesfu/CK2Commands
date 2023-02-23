package urzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克札尔AkzharBarony struct {
	feud.BaseBarony
}

var BAkzhar阿克札尔 feud.Barony = &阿克札尔AkzharBarony{}

func init() {
	f := BAkzhar阿克札尔.(*阿克札尔AkzharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akzhar",
		TitleName: "阿克札尔",
		TitleCode: "b_akzhar",
	}
}
