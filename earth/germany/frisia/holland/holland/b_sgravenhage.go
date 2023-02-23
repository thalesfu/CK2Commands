package holland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海牙SgravenhageBarony struct {
	feud.BaseBarony
}

var BSgravenhage海牙 feud.Barony = &海牙SgravenhageBarony{}

func init() {
	f := BSgravenhage海牙.(*海牙SgravenhageBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sgravenhage",
		TitleName: "海牙",
		TitleCode: "b_sgravenhage",
	}
}
