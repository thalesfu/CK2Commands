package nandana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欢喜园NandanaBarony struct {
	feud.BaseBarony
}

var BNandana欢喜园 feud.Barony = &欢喜园NandanaBarony{}

func init() {
	f := BNandana欢喜园.(*欢喜园NandanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nandana",
		TitleName: "欢喜园",
		TitleCode: "b_nandana",
	}
}
