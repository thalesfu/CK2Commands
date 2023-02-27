package csanad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍德迈泽瓦沙尔海伊HodmezovasarhelyBarony struct {
	feud.BaseBarony
}

var BHodmezovasarhely霍德迈泽瓦沙尔海伊 feud.Barony = &霍德迈泽瓦沙尔海伊HodmezovasarhelyBarony{}

func init() {
    f := BHodmezovasarhely霍德迈泽瓦沙尔海伊.(*霍德迈泽瓦沙尔海伊HodmezovasarhelyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hodmezovasarhely",
		TitleName: "霍德迈泽瓦沙尔海伊",
		TitleCode: "b_hodmezovasarhely",
	}
}
