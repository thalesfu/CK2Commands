package imeretia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库塔伊西KutaisiBarony struct {
	feud.BaseBarony
}

var BKutaisi库塔伊西 feud.Barony = &库塔伊西KutaisiBarony{}

func init() {
	f := BKutaisi库塔伊西.(*库塔伊西KutaisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kutaisi",
		TitleName: "库塔伊西",
		TitleCode: "b_kutaisi",
	}
}
