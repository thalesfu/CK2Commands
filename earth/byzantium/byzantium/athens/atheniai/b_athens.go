package atheniai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅典AthensBarony struct {
	feud.BaseBarony
}

var BAthens雅典 feud.Barony = &雅典AthensBarony{}

func init() {
    f := BAthens雅典.(*雅典AthensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "athens",
		TitleName: "雅典",
		TitleCode: "b_athens",
	}
}
