package huelva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔哈拉克AljaraqueBarony struct {
	feud.BaseBarony
}

var BAljaraque阿尔哈拉克 feud.Barony = &阿尔哈拉克AljaraqueBarony{}

func init() {
	f := BAljaraque阿尔哈拉克.(*阿尔哈拉克AljaraqueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aljaraque",
		TitleName: "阿尔哈拉克",
		TitleCode: "b_aljaraque",
	}
}
