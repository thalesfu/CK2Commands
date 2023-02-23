package tom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绍罗霍沃ShorokhovoBarony struct {
	feud.BaseBarony
}

var BShorokhovo绍罗霍沃 feud.Barony = &绍罗霍沃ShorokhovoBarony{}

func init() {
	f := BShorokhovo绍罗霍沃.(*绍罗霍沃ShorokhovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shorokhovo",
		TitleName: "绍罗霍沃",
		TitleCode: "b_shorokhovo",
	}
}
