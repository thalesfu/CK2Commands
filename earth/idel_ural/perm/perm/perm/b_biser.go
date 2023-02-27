package perm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比谢尔BiserBarony struct {
	feud.BaseBarony
}

var BBiser比谢尔 feud.Barony = &比谢尔BiserBarony{}

func init() {
    f := BBiser比谢尔.(*比谢尔BiserBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biser",
		TitleName: "比谢尔",
		TitleCode: "b_biser",
	}
}
