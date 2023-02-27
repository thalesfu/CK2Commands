package balkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 忽懔KhulmBarony struct {
	feud.BaseBarony
}

var BKhulm忽懔 feud.Barony = &忽懔KhulmBarony{}

func init() {
    f := BKhulm忽懔.(*忽懔KhulmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khulm",
		TitleName: "忽懔",
		TitleCode: "b_khulm",
	}
}
