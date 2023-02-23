package senlis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡比涅CompiegneBarony struct {
	feud.BaseBarony
}

var BCompiegne贡比涅 feud.Barony = &贡比涅CompiegneBarony{}

func init() {
	f := BCompiegne贡比涅.(*贡比涅CompiegneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "compiegne",
		TitleName: "贡比涅",
		TitleCode: "b_compiegne",
	}
}
