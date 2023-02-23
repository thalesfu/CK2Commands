package kabul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库纳尔KunarBarony struct {
	feud.BaseBarony
}

var BKunar库纳尔 feud.Barony = &库纳尔KunarBarony{}

func init() {
	f := BKunar库纳尔.(*库纳尔KunarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunar",
		TitleName: "库纳尔",
		TitleCode: "b_kunar",
	}
}
