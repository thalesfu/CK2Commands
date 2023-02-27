package naldurg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 槃尼耶伽罗摩BaniagaonBarony struct {
	feud.BaseBarony
}

var BBaniagaon槃尼耶伽罗摩 feud.Barony = &槃尼耶伽罗摩BaniagaonBarony{}

func init() {
    f := BBaniagaon槃尼耶伽罗摩.(*槃尼耶伽罗摩BaniagaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baniagaon",
		TitleName: "槃尼耶伽罗摩",
		TitleCode: "b_baniagaon",
	}
}
