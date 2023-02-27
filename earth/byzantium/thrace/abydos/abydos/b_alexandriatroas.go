package abydos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚历山大特罗亚AlexandriatroasBarony struct {
	feud.BaseBarony
}

var BAlexandriatroas亚历山大特罗亚 feud.Barony = &亚历山大特罗亚AlexandriatroasBarony{}

func init() {
    f := BAlexandriatroas亚历山大特罗亚.(*亚历山大特罗亚AlexandriatroasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alexandriatroas",
		TitleName: "亚历山大特罗亚",
		TitleCode: "b_alexandriatroas",
	}
}
