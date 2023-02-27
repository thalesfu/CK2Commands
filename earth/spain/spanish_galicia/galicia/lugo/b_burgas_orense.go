package lugo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔加斯Burgas_orenseBarony struct {
	feud.BaseBarony
}

var BBurgas_orense布尔加斯 feud.Barony = &布尔加斯Burgas_orenseBarony{}

func init() {
    f := BBurgas_orense布尔加斯.(*布尔加斯Burgas_orenseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burgas_orense",
		TitleName: "布尔加斯",
		TitleCode: "b_burgas_orense",
	}
}
