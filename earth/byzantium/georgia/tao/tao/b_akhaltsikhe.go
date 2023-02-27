package tao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿哈尔齐赫AkhaltsikheBarony struct {
	feud.BaseBarony
}

var BAkhaltsikhe阿哈尔齐赫 feud.Barony = &阿哈尔齐赫AkhaltsikheBarony{}

func init() {
    f := BAkhaltsikhe阿哈尔齐赫.(*阿哈尔齐赫AkhaltsikheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akhaltsikhe",
		TitleName: "阿哈尔齐赫",
		TitleCode: "b_akhaltsikhe",
	}
}
