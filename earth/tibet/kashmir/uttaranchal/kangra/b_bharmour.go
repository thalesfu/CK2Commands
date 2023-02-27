package kangra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆罗吸摩补罗BharmourBarony struct {
	feud.BaseBarony
}

var BBharmour婆罗吸摩补罗 feud.Barony = &婆罗吸摩补罗BharmourBarony{}

func init() {
    f := BBharmour婆罗吸摩补罗.(*婆罗吸摩补罗BharmourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bharmour",
		TitleName: "婆罗吸摩补罗",
		TitleCode: "b_bharmour",
	}
}
