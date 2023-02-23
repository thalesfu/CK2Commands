package middlesex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富勒姆FulhamBarony struct {
	feud.BaseBarony
}

var BFulham富勒姆 feud.Barony = &富勒姆FulhamBarony{}

func init() {
	f := BFulham富勒姆.(*富勒姆FulhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fulham",
		TitleName: "富勒姆",
		TitleCode: "b_fulham",
	}
}
