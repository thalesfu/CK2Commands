package venadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳盖科伊尔NagercoilBarony struct {
	feud.BaseBarony
}

var BNagercoil纳盖科伊尔 feud.Barony = &纳盖科伊尔NagercoilBarony{}

func init() {
    f := BNagercoil纳盖科伊尔.(*纳盖科伊尔NagercoilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagercoil",
		TitleName: "纳盖科伊尔",
		TitleCode: "b_nagercoil",
	}
}
