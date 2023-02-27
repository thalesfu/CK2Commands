package besancon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝桑松BesanconBarony struct {
	feud.BaseBarony
}

var BBesancon贝桑松 feud.Barony = &贝桑松BesanconBarony{}

func init() {
    f := BBesancon贝桑松.(*贝桑松BesanconBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "besancon",
		TitleName: "贝桑松",
		TitleCode: "b_besancon",
	}
}
