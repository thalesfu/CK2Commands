package anglesey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本莱赫BenllechBarony struct {
	feud.BaseBarony
}

var BBenllech本莱赫 feud.Barony = &本莱赫BenllechBarony{}

func init() {
    f := BBenllech本莱赫.(*本莱赫BenllechBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "benllech",
		TitleName: "本莱赫",
		TitleCode: "b_benllech",
	}
}
