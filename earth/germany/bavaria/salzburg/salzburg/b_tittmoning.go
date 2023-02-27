package salzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂特莫宁TittmoningBarony struct {
	feud.BaseBarony
}

var BTittmoning蒂特莫宁 feud.Barony = &蒂特莫宁TittmoningBarony{}

func init() {
    f := BTittmoning蒂特莫宁.(*蒂特莫宁TittmoningBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tittmoning",
		TitleName: "蒂特莫宁",
		TitleCode: "b_tittmoning",
	}
}
