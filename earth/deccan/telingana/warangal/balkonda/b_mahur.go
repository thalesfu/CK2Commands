package balkonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛赫尔MahurBarony struct {
	feud.BaseBarony
}

var BMahur玛赫尔 feud.Barony = &玛赫尔MahurBarony{}

func init() {
    f := BMahur玛赫尔.(*玛赫尔MahurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahur",
		TitleName: "玛赫尔",
		TitleCode: "b_mahur",
	}
}
