package koln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波恩BonnBarony struct {
	feud.BaseBarony
}

var BBonn波恩 feud.Barony = &波恩BonnBarony{}

func init() {
    f := BBonn波恩.(*波恩BonnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bonn",
		TitleName: "波恩",
		TitleCode: "b_bonn",
	}
}
