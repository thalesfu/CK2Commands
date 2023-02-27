package kastamon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿柏农AbononBarony struct {
	feud.BaseBarony
}

var BAbonon阿柏农 feud.Barony = &阿柏农AbononBarony{}

func init() {
    f := BAbonon阿柏农.(*阿柏农AbononBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abonon",
		TitleName: "阿柏农",
		TitleCode: "b_abonon",
	}
}
