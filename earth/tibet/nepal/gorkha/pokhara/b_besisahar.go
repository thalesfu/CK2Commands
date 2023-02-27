package pokhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝西沙尔BesisaharBarony struct {
	feud.BaseBarony
}

var BBesisahar贝西沙尔 feud.Barony = &贝西沙尔BesisaharBarony{}

func init() {
    f := BBesisahar贝西沙尔.(*贝西沙尔BesisaharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "besisahar",
		TitleName: "贝西沙尔",
		TitleCode: "b_besisahar",
	}
}
