package kholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲绍沃FishovoBarony struct {
	feud.BaseBarony
}

var BFishovo菲绍沃 feud.Barony = &菲绍沃FishovoBarony{}

func init() {
	f := BFishovo菲绍沃.(*菲绍沃FishovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fishovo",
		TitleName: "菲绍沃",
		TitleCode: "b_fishovo",
	}
}
