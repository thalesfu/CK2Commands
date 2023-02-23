package kakheti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹韦利加拉瓦尼DzveligalavaniBarony struct {
	feud.BaseBarony
}

var BDzveligalavani兹韦利加拉瓦尼 feud.Barony = &兹韦利加拉瓦尼DzveligalavaniBarony{}

func init() {
	f := BDzveligalavani兹韦利加拉瓦尼.(*兹韦利加拉瓦尼DzveligalavaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dzveligalavani",
		TitleName: "兹韦利加拉瓦尼",
		TitleCode: "b_dzveligalavani",
	}
}
