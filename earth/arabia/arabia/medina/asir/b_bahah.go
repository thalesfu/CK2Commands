package asir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴哈BahahBarony struct {
	feud.BaseBarony
}

var BBahah巴哈 feud.Barony = &巴哈BahahBarony{}

func init() {
    f := BBahah巴哈.(*巴哈BahahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bahah",
		TitleName: "巴哈",
		TitleCode: "b_bahah",
	}
}
