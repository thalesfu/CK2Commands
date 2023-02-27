package gyantse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康马KangmarBarony struct {
	feud.BaseBarony
}

var BKangmar康马 feud.Barony = &康马KangmarBarony{}

func init() {
    f := BKangmar康马.(*康马KangmarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kangmar",
		TitleName: "康马",
		TitleCode: "b_kangmar",
	}
}
