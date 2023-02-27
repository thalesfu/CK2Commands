package brie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普罗万ProvinsBarony struct {
	feud.BaseBarony
}

var BProvins普罗万 feud.Barony = &普罗万ProvinsBarony{}

func init() {
    f := BProvins普罗万.(*普罗万ProvinsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "provins",
		TitleName: "普罗万",
		TitleCode: "b_provins",
	}
}
