package lubeck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施卢图普SchlutupBarony struct {
	feud.BaseBarony
}

var BSchlutup施卢图普 feud.Barony = &施卢图普SchlutupBarony{}

func init() {
    f := BSchlutup施卢图普.(*施卢图普SchlutupBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schlutup",
		TitleName: "施卢图普",
		TitleCode: "b_schlutup",
	}
}
