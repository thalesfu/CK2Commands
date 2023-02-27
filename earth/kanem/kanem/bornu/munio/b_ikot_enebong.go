package munio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊科特埃内邦Ikot_enebongBarony struct {
	feud.BaseBarony
}

var BIkot_enebong伊科特埃内邦 feud.Barony = &伊科特埃内邦Ikot_enebongBarony{}

func init() {
    f := BIkot_enebong伊科特埃内邦.(*伊科特埃内邦Ikot_enebongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ikot_enebong",
		TitleName: "伊科特埃内邦",
		TitleCode: "b_ikot_enebong",
	}
}
