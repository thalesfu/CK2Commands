package dhu_zabi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀图尔QutufBarony struct {
	feud.BaseBarony
}

var BQutuf喀图尔 feud.Barony = &喀图尔QutufBarony{}

func init() {
    f := BQutuf喀图尔.(*喀图尔QutufBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qutuf",
		TitleName: "喀图尔",
		TitleCode: "b_qutuf",
	}
}
