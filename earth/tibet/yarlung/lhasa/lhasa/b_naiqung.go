package lhasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乃琼NaiqungBarony struct {
	feud.BaseBarony
}

var BNaiqung乃琼 feud.Barony = &乃琼NaiqungBarony{}

func init() {
    f := BNaiqung乃琼.(*乃琼NaiqungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naiqung",
		TitleName: "乃琼",
		TitleCode: "b_naiqung",
	}
}
