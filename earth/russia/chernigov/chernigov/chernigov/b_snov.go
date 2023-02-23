package chernigov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯诺夫SnovBarony struct {
	feud.BaseBarony
}

var BSnov斯诺夫 feud.Barony = &斯诺夫SnovBarony{}

func init() {
	f := BSnov斯诺夫.(*斯诺夫SnovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "snov",
		TitleName: "斯诺夫",
		TitleCode: "b_snov",
	}
}
