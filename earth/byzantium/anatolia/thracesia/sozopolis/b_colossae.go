package sozopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科洛塞ColossaeBarony struct {
	feud.BaseBarony
}

var BColossae科洛塞 feud.Barony = &科洛塞ColossaeBarony{}

func init() {
    f := BColossae科洛塞.(*科洛塞ColossaeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "colossae",
		TitleName: "科洛塞",
		TitleCode: "b_colossae",
	}
}
