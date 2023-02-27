package ludrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾沙梅尔JaisalmerBarony struct {
	feud.BaseBarony
}

var BJaisalmer贾沙梅尔 feud.Barony = &贾沙梅尔JaisalmerBarony{}

func init() {
    f := BJaisalmer贾沙梅尔.(*贾沙梅尔JaisalmerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaisalmer",
		TitleName: "贾沙梅尔",
		TitleCode: "b_jaisalmer",
	}
}
