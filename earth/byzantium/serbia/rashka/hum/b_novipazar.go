package hum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔戈维什泰NovipazarBarony struct {
	feud.BaseBarony
}

var BNovipazar特尔戈维什泰 feud.Barony = &特尔戈维什泰NovipazarBarony{}

func init() {
	f := BNovipazar特尔戈维什泰.(*特尔戈维什泰NovipazarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novipazar",
		TitleName: "特尔戈维什泰",
		TitleCode: "b_novipazar",
	}
}
