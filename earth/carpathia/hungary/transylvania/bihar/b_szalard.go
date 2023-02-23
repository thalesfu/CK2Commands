package bihar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉尔德SzalardBarony struct {
	feud.BaseBarony
}

var BSzalard萨拉尔德 feud.Barony = &萨拉尔德SzalardBarony{}

func init() {
	f := BSzalard萨拉尔德.(*萨拉尔德SzalardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szalard",
		TitleName: "萨拉尔德",
		TitleCode: "b_szalard",
	}
}
