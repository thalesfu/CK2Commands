package hellas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱瓦贾LevadhiaBarony struct {
	feud.BaseBarony
}

var BLevadhia莱瓦贾 feud.Barony = &莱瓦贾LevadhiaBarony{}

func init() {
    f := BLevadhia莱瓦贾.(*莱瓦贾LevadhiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "levadhia",
		TitleName: "莱瓦贾",
		TitleCode: "b_levadhia",
	}
}
