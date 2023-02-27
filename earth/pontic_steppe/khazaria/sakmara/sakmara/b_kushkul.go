package sakmara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库什库尔KushkulBarony struct {
	feud.BaseBarony
}

var BKushkul库什库尔 feud.Barony = &库什库尔KushkulBarony{}

func init() {
    f := BKushkul库什库尔.(*库什库尔KushkulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kushkul",
		TitleName: "库什库尔",
		TitleCode: "b_kushkul",
	}
}
