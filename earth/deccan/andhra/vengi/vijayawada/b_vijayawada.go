package vijayawada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗阇耶婆陀VijayawadaBarony struct {
	feud.BaseBarony
}

var BVijayawada毗阇耶婆陀 feud.Barony = &毗阇耶婆陀VijayawadaBarony{}

func init() {
	f := BVijayawada毗阇耶婆陀.(*毗阇耶婆陀VijayawadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vijayawada",
		TitleName: "毗阇耶婆陀",
		TitleCode: "b_vijayawada",
	}
}
