package galaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝雷什蒂BerestiBarony struct {
	feud.BaseBarony
}

var BBeresti贝雷什蒂 feud.Barony = &贝雷什蒂BerestiBarony{}

func init() {
    f := BBeresti贝雷什蒂.(*贝雷什蒂BerestiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beresti",
		TitleName: "贝雷什蒂",
		TitleCode: "b_beresti",
	}
}
