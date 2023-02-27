package belukha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别卢哈BelukhaBarony struct {
	feud.BaseBarony
}

var BBelukha别卢哈 feud.Barony = &别卢哈BelukhaBarony{}

func init() {
    f := BBelukha别卢哈.(*别卢哈BelukhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belukha",
		TitleName: "别卢哈",
		TitleCode: "b_belukha",
	}
}
