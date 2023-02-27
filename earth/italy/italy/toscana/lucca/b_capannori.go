package lucca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡潘诺里CapannoriBarony struct {
	feud.BaseBarony
}

var BCapannori卡潘诺里 feud.Barony = &卡潘诺里CapannoriBarony{}

func init() {
    f := BCapannori卡潘诺里.(*卡潘诺里CapannoriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "capannori",
		TitleName: "卡潘诺里",
		TitleCode: "b_capannori",
	}
}
