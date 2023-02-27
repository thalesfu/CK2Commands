package trencin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班BanBarony struct {
	feud.BaseBarony
}

var BBan班 feud.Barony = &班BanBarony{}

func init() {
    f := BBan班.(*班BanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ban",
		TitleName: "班",
		TitleCode: "b_ban",
	}
}
