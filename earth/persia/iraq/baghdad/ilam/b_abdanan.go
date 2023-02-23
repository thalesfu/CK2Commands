package ilam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卜达南AbdananBarony struct {
	feud.BaseBarony
}

var BAbdanan阿卜达南 feud.Barony = &阿卜达南AbdananBarony{}

func init() {
	f := BAbdanan阿卜达南.(*阿卜达南AbdananBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abdanan",
		TitleName: "阿卜达南",
		TitleCode: "b_abdanan",
	}
}
