package uchangidurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尸罗SiraBarony struct {
	feud.BaseBarony
}

var BSira尸罗 feud.Barony = &尸罗SiraBarony{}

func init() {
    f := BSira尸罗.(*尸罗SiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sira",
		TitleName: "尸罗",
		TitleCode: "b_sira",
	}
}
