package mahoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃弥罗城HamirpurBarony struct {
	feud.BaseBarony
}

var BHamirpur诃弥罗城 feud.Barony = &诃弥罗城HamirpurBarony{}

func init() {
    f := BHamirpur诃弥罗城.(*诃弥罗城HamirpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamirpur",
		TitleName: "诃弥罗城",
		TitleCode: "b_hamirpur",
	}
}
