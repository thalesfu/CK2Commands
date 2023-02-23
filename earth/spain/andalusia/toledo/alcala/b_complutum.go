package alcala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔普卢图姆ComplutumBarony struct {
	feud.BaseBarony
}

var BComplutum孔普卢图姆 feud.Barony = &孔普卢图姆ComplutumBarony{}

func init() {
	f := BComplutum孔普卢图姆.(*孔普卢图姆ComplutumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "complutum",
		TitleName: "孔普卢图姆",
		TitleCode: "b_complutum",
	}
}
