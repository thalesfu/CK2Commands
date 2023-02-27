package perm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩雷姆PerymBarony struct {
	feud.BaseBarony
}

var BPerym佩雷姆 feud.Barony = &佩雷姆PerymBarony{}

func init() {
    f := BPerym佩雷姆.(*佩雷姆PerymBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perym",
		TitleName: "佩雷姆",
		TitleCode: "b_perym",
	}
}
