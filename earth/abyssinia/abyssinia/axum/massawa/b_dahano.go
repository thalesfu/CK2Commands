package massawa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达哈努DahanoBarony struct {
	feud.BaseBarony
}

var BDahano达哈努 feud.Barony = &达哈努DahanoBarony{}

func init() {
	f := BDahano达哈努.(*达哈努DahanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dahano",
		TitleName: "达哈努",
		TitleCode: "b_dahano",
	}
}
