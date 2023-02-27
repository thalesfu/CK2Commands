package tuul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳来哈NalaikhBarony struct {
	feud.BaseBarony
}

var BNalaikh纳来哈 feud.Barony = &纳来哈NalaikhBarony{}

func init() {
    f := BNalaikh纳来哈.(*纳来哈NalaikhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nalaikh",
		TitleName: "纳来哈",
		TitleCode: "b_nalaikh",
	}
}
