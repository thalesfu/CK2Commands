package yumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下西号XiaxihaoBarony struct {
	feud.BaseBarony
}

var BXiaxihao下西号 feud.Barony = &下西号XiaxihaoBarony{}

func init() {
    f := BXiaxihao下西号.(*下西号XiaxihaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xiaxihao",
		TitleName: "下西号",
		TitleCode: "b_xiaxihao",
	}
}
