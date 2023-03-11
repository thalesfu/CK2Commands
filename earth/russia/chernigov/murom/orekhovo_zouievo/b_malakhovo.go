package orekhovo_zouievo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉霍沃MalakhovoBarony struct {
	feud.BaseBarony
}

var BMalakhovo马拉霍沃 feud.Barony = &马拉霍沃MalakhovoBarony{}

func init() {
    f := BMalakhovo马拉霍沃.(*马拉霍沃MalakhovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malakhovo",
		TitleName: "马拉霍沃",
		TitleCode: "b_malakhovo",
	}
}
