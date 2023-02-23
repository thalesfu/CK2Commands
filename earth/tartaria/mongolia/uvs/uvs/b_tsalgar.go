package uvs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察勒噶尔TsalgarBarony struct {
	feud.BaseBarony
}

var BTsalgar察勒噶尔 feud.Barony = &察勒噶尔TsalgarBarony{}

func init() {
	f := BTsalgar察勒噶尔.(*察勒噶尔TsalgarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsalgar",
		TitleName: "察勒噶尔",
		TitleCode: "b_tsalgar",
	}
}
