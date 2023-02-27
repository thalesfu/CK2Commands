package pitten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格洛格尼茨GloggnitzBarony struct {
	feud.BaseBarony
}

var BGloggnitz格洛格尼茨 feud.Barony = &格洛格尼茨GloggnitzBarony{}

func init() {
    f := BGloggnitz格洛格尼茨.(*格洛格尼茨GloggnitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gloggnitz",
		TitleName: "格洛格尼茨",
		TitleCode: "b_gloggnitz",
	}
}
