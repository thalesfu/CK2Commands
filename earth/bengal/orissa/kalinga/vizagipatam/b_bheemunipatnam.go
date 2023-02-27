package vizagipatam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠夷牟尼跋南BheemunipatnamBarony struct {
	feud.BaseBarony
}

var BBheemunipatnam吠夷牟尼跋南 feud.Barony = &吠夷牟尼跋南BheemunipatnamBarony{}

func init() {
    f := BBheemunipatnam吠夷牟尼跋南.(*吠夷牟尼跋南BheemunipatnamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bheemunipatnam",
		TitleName: "吠夷牟尼跋南",
		TitleCode: "b_bheemunipatnam",
	}
}
