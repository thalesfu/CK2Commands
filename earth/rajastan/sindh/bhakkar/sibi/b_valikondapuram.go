package sibi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伐梨军荼补蓝ValikondapuramBarony struct {
	feud.BaseBarony
}

var BValikondapuram伐梨军荼补蓝 feud.Barony = &伐梨军荼补蓝ValikondapuramBarony{}

func init() {
	f := BValikondapuram伐梨军荼补蓝.(*伐梨军荼补蓝ValikondapuramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valikondapuram",
		TitleName: "伐梨军荼补蓝",
		TitleCode: "b_valikondapuram",
	}
}
