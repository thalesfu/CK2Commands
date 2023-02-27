package mountain_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛比尔斯克SimbirskBarony struct {
	feud.BaseBarony
}

var BSimbirsk辛比尔斯克 feud.Barony = &辛比尔斯克SimbirskBarony{}

func init() {
    f := BSimbirsk辛比尔斯克.(*辛比尔斯克SimbirskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "simbirsk",
		TitleName: "辛比尔斯克",
		TitleCode: "b_simbirsk",
	}
}
