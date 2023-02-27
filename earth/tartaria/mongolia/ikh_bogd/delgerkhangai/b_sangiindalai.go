package delgerkhangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑根达来SangiindalaiBarony struct {
	feud.BaseBarony
}

var BSangiindalai桑根达来 feud.Barony = &桑根达来SangiindalaiBarony{}

func init() {
    f := BSangiindalai桑根达来.(*桑根达来SangiindalaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangiindalai",
		TitleName: "桑根达来",
		TitleCode: "b_sangiindalai",
	}
}
