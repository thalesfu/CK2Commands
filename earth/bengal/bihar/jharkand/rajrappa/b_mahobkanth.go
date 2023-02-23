package rajrappa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩呼乾MahobkanthBarony struct {
	feud.BaseBarony
}

var BMahobkanth摩呼乾 feud.Barony = &摩呼乾MahobkanthBarony{}

func init() {
	f := BMahobkanth摩呼乾.(*摩呼乾MahobkanthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahobkanth",
		TitleName: "摩呼乾",
		TitleCode: "b_mahobkanth",
	}
}
