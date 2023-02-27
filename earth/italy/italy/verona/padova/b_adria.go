package padova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德里亚AdriaBarony struct {
	feud.BaseBarony
}

var BAdria阿德里亚 feud.Barony = &阿德里亚AdriaBarony{}

func init() {
    f := BAdria阿德里亚.(*阿德里亚AdriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adria",
		TitleName: "阿德里亚",
		TitleCode: "b_adria",
	}
}
