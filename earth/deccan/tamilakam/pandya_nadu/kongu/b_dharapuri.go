package kongu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀罗补罗DharapuriBarony struct {
	feud.BaseBarony
}

var BDharapuri陀罗补罗 feud.Barony = &陀罗补罗DharapuriBarony{}

func init() {
    f := BDharapuri陀罗补罗.(*陀罗补罗DharapuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dharapuri",
		TitleName: "陀罗补罗",
		TitleCode: "b_dharapuri",
	}
}
