package slupsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔加德BelgardBarony struct {
	feud.BaseBarony
}

var BBelgard贝尔加德 feud.Barony = &贝尔加德BelgardBarony{}

func init() {
    f := BBelgard贝尔加德.(*贝尔加德BelgardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belgard",
		TitleName: "贝尔加德",
		TitleCode: "b_belgard",
	}
}
