package lappland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴依瓦特奈BergvattnetBarony struct {
	feud.BaseBarony
}

var BBergvattnet巴依瓦特奈 feud.Barony = &巴依瓦特奈BergvattnetBarony{}

func init() {
    f := BBergvattnet巴依瓦特奈.(*巴依瓦特奈BergvattnetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bergvattnet",
		TitleName: "巴依瓦特奈",
		TitleCode: "b_bergvattnet",
	}
}
