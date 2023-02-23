package fergana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳达克NadaqBarony struct {
	feud.BaseBarony
}

var BNadaq纳达克 feud.Barony = &纳达克NadaqBarony{}

func init() {
	f := BNadaq纳达克.(*纳达克NadaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nadaq",
		TitleName: "纳达克",
		TitleCode: "b_nadaq",
	}
}
