package sanaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾拜迪亚JabaltiyalBarony struct {
	feud.BaseBarony
}

var BJabaltiyal贾拜迪亚 feud.Barony = &贾拜迪亚JabaltiyalBarony{}

func init() {
    f := BJabaltiyal贾拜迪亚.(*贾拜迪亚JabaltiyalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jabaltiyal",
		TitleName: "贾拜迪亚",
		TitleCode: "b_jabaltiyal",
	}
}
