package tannu_ola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏特霍尔Sut_kholBarony struct {
	feud.BaseBarony
}

var BSut_khol苏特霍尔 feud.Barony = &苏特霍尔Sut_kholBarony{}

func init() {
    f := BSut_khol苏特霍尔.(*苏特霍尔Sut_kholBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sut_khol",
		TitleName: "苏特霍尔",
		TitleCode: "b_sut_khol",
	}
}
