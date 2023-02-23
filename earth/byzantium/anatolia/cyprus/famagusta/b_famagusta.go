package famagusta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法马古斯塔FamagustaBarony struct {
	feud.BaseBarony
}

var BFamagusta法马古斯塔 feud.Barony = &法马古斯塔FamagustaBarony{}

func init() {
	f := BFamagusta法马古斯塔.(*法马古斯塔FamagustaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "famagusta",
		TitleName: "法马古斯塔",
		TitleCode: "b_famagusta",
	}
}
