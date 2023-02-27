package kolomna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉普捷沃LaptevoBarony struct {
	feud.BaseBarony
}

var BLaptevo拉普捷沃 feud.Barony = &拉普捷沃LaptevoBarony{}

func init() {
    f := BLaptevo拉普捷沃.(*拉普捷沃LaptevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laptevo",
		TitleName: "拉普捷沃",
		TitleCode: "b_laptevo",
	}
}
