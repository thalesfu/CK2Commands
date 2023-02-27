package euphrates

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法里斯FarisBarony struct {
	feud.BaseBarony
}

var BFaris法里斯 feud.Barony = &法里斯FarisBarony{}

func init() {
    f := BFaris法里斯.(*法里斯FarisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faris",
		TitleName: "法里斯",
		TitleCode: "b_faris",
	}
}
