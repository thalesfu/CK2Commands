package sarasvata_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍陀SherthaBarony struct {
	feud.BaseBarony
}

var BShertha舍陀 feud.Barony = &舍陀SherthaBarony{}

func init() {
    f := BShertha舍陀.(*舍陀SherthaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shertha",
		TitleName: "舍陀",
		TitleCode: "b_shertha",
	}
}
