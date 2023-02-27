package chikoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切列姆霍沃CheremkhovoBarony struct {
	feud.BaseBarony
}

var BCheremkhovo切列姆霍沃 feud.Barony = &切列姆霍沃CheremkhovoBarony{}

func init() {
    f := BCheremkhovo切列姆霍沃.(*切列姆霍沃CheremkhovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cheremkhovo",
		TitleName: "切列姆霍沃",
		TitleCode: "b_cheremkhovo",
	}
}
