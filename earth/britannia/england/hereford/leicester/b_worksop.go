package leicester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃克索普WorksopBarony struct {
	feud.BaseBarony
}

var BWorksop沃克索普 feud.Barony = &沃克索普WorksopBarony{}

func init() {
    f := BWorksop沃克索普.(*沃克索普WorksopBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "worksop",
		TitleName: "沃克索普",
		TitleCode: "b_worksop",
	}
}
