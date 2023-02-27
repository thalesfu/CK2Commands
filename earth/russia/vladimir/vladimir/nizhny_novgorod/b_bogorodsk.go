package nizhny_novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博戈罗茨克BogorodskBarony struct {
	feud.BaseBarony
}

var BBogorodsk博戈罗茨克 feud.Barony = &博戈罗茨克BogorodskBarony{}

func init() {
    f := BBogorodsk博戈罗茨克.(*博戈罗茨克BogorodskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bogorodsk",
		TitleName: "博戈罗茨克",
		TitleCode: "b_bogorodsk",
	}
}
