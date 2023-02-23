package vijnot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般提耶BhatiyaBarony struct {
	feud.BaseBarony
}

var BBhatiya般提耶 feud.Barony = &般提耶BhatiyaBarony{}

func init() {
	f := BBhatiya般提耶.(*般提耶BhatiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhatiya",
		TitleName: "般提耶",
		TitleCode: "b_bhatiya",
	}
}
