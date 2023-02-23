package naupactus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿吕齐亚AlyzeiaBarony struct {
	feud.BaseBarony
}

var BAlyzeia阿吕齐亚 feud.Barony = &阿吕齐亚AlyzeiaBarony{}

func init() {
	f := BAlyzeia阿吕齐亚.(*阿吕齐亚AlyzeiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alyzeia",
		TitleName: "阿吕齐亚",
		TitleCode: "b_alyzeia",
	}
}
