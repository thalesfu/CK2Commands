package acalapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 槃瞿闷荼BangomundaBarony struct {
	feud.BaseBarony
}

var BBangomunda槃瞿闷荼 feud.Barony = &槃瞿闷荼BangomundaBarony{}

func init() {
    f := BBangomunda槃瞿闷荼.(*槃瞿闷荼BangomundaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bangomunda",
		TitleName: "槃瞿闷荼",
		TitleCode: "b_bangomunda",
	}
}
