package gilgit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弥那伐尔MinawarBarony struct {
	feud.BaseBarony
}

var BMinawar弥那伐尔 feud.Barony = &弥那伐尔MinawarBarony{}

func init() {
	f := BMinawar弥那伐尔.(*弥那伐尔MinawarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "minawar",
		TitleName: "弥那伐尔",
		TitleCode: "b_minawar",
	}
}
