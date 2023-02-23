package farama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟燕SeyanBarony struct {
	feud.BaseBarony
}

var BSeyan瑟燕 feud.Barony = &瑟燕SeyanBarony{}

func init() {
	f := BSeyan瑟燕.(*瑟燕SeyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seyan",
		TitleName: "瑟燕",
		TitleCode: "b_seyan",
	}
}
