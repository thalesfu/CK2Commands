package alampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗阇耶那揭罗VijayanagaraBarony struct {
	feud.BaseBarony
}

var BVijayanagara毗阇耶那揭罗 feud.Barony = &毗阇耶那揭罗VijayanagaraBarony{}

func init() {
	f := BVijayanagara毗阇耶那揭罗.(*毗阇耶那揭罗VijayanagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vijayanagara",
		TitleName: "毗阇耶那揭罗",
		TitleCode: "b_vijayanagara",
	}
}
