package friesach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼得斯贝格PetersbergBarony struct {
	feud.BaseBarony
}

var BPetersberg彼得斯贝格 feud.Barony = &彼得斯贝格PetersbergBarony{}

func init() {
	f := BPetersberg彼得斯贝格.(*彼得斯贝格PetersbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petersberg",
		TitleName: "彼得斯贝格",
		TitleCode: "b_petersberg",
	}
}
