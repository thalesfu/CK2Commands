package reval

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈普萨卢HapsalBarony struct {
	feud.BaseBarony
}

var BHapsal哈普萨卢 feud.Barony = &哈普萨卢HapsalBarony{}

func init() {
	f := BHapsal哈普萨卢.(*哈普萨卢HapsalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hapsal",
		TitleName: "哈普萨卢",
		TitleCode: "b_hapsal",
	}
}
