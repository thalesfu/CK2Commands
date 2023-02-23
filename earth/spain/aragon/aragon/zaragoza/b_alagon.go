package zaragoza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉贡AlagonBarony struct {
	feud.BaseBarony
}

var BAlagon阿拉贡 feud.Barony = &阿拉贡AlagonBarony{}

func init() {
	f := BAlagon阿拉贡.(*阿拉贡AlagonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alagon",
		TitleName: "阿拉贡",
		TitleCode: "b_alagon",
	}
}
