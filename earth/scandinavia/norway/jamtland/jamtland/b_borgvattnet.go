package jamtland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博里瓦特讷BorgvattnetBarony struct {
	feud.BaseBarony
}

var BBorgvattnet博里瓦特讷 feud.Barony = &博里瓦特讷BorgvattnetBarony{}

func init() {
    f := BBorgvattnet博里瓦特讷.(*博里瓦特讷BorgvattnetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borgvattnet",
		TitleName: "博里瓦特讷",
		TitleCode: "b_borgvattnet",
	}
}
