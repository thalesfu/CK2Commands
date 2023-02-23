package dihistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔沙克利TorskhaliBarony struct {
	feud.BaseBarony
}

var BTorskhali图尔沙克利 feud.Barony = &图尔沙克利TorskhaliBarony{}

func init() {
	f := BTorskhali图尔沙克利.(*图尔沙克利TorskhaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torskhali",
		TitleName: "图尔沙克利",
		TitleCode: "b_torskhali",
	}
}
