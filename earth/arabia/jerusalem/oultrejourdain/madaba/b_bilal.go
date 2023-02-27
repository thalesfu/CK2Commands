package madaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比拉勒BilalBarony struct {
	feud.BaseBarony
}

var BBilal比拉勒 feud.Barony = &比拉勒BilalBarony{}

func init() {
    f := BBilal比拉勒.(*比拉勒BilalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilal",
		TitleName: "比拉勒",
		TitleCode: "b_bilal",
	}
}
