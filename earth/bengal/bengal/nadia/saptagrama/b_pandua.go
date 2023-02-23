package saptagrama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般菟阿PanduaBarony struct {
	feud.BaseBarony
}

var BPandua般菟阿 feud.Barony = &般菟阿PanduaBarony{}

func init() {
	f := BPandua般菟阿.(*般菟阿PanduaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pandua",
		TitleName: "般菟阿",
		TitleCode: "b_pandua",
	}
}
