package orava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔沃堡ArvavaraBarony struct {
	feud.BaseBarony
}

var BArvavara阿尔沃堡 feud.Barony = &阿尔沃堡ArvavaraBarony{}

func init() {
	f := BArvavara阿尔沃堡.(*阿尔沃堡ArvavaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arvavara",
		TitleName: "阿尔沃堡",
		TitleCode: "b_arvavara",
	}
}
