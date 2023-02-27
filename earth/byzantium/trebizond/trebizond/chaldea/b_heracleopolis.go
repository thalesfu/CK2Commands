package chaldea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫拉克莱奥波利斯HeracleopolisBarony struct {
	feud.BaseBarony
}

var BHeracleopolis赫拉克莱奥波利斯 feud.Barony = &赫拉克莱奥波利斯HeracleopolisBarony{}

func init() {
    f := BHeracleopolis赫拉克莱奥波利斯.(*赫拉克莱奥波利斯HeracleopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heracleopolis",
		TitleName: "赫拉克莱奥波利斯",
		TitleCode: "b_heracleopolis",
	}
}
