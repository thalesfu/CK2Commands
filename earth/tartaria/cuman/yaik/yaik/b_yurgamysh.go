package yaik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤尔加梅什YurgamyshBarony struct {
	feud.BaseBarony
}

var BYurgamysh尤尔加梅什 feud.Barony = &尤尔加梅什YurgamyshBarony{}

func init() {
	f := BYurgamysh尤尔加梅什.(*尤尔加梅什YurgamyshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yurgamysh",
		TitleName: "尤尔加梅什",
		TitleCode: "b_yurgamysh",
	}
}
