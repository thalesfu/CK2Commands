package aprutium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩内Aprutium_penneBarony struct {
	feud.BaseBarony
}

var BAprutium_penne佩内 feud.Barony = &佩内Aprutium_penneBarony{}

func init() {
    f := BAprutium_penne佩内.(*佩内Aprutium_penneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aprutium_penne",
		TitleName: "佩内",
		TitleCode: "b_aprutium_penne",
	}
}
