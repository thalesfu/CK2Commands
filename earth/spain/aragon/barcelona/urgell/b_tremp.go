package urgell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特伦普TrempBarony struct {
	feud.BaseBarony
}

var BTremp特伦普 feud.Barony = &特伦普TrempBarony{}

func init() {
    f := BTremp特伦普.(*特伦普TrempBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tremp",
		TitleName: "特伦普",
		TitleCode: "b_tremp",
	}
}
