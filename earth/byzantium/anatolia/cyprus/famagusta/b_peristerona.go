package famagusta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 派里斯泰罗纳PeristeronaBarony struct {
	feud.BaseBarony
}

var BPeristerona派里斯泰罗纳 feud.Barony = &派里斯泰罗纳PeristeronaBarony{}

func init() {
    f := BPeristerona派里斯泰罗纳.(*派里斯泰罗纳PeristeronaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peristerona",
		TitleName: "派里斯泰罗纳",
		TitleCode: "b_peristerona",
	}
}
