package atlas_mnt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈西本马贾HassibenmadjebBarony struct {
	feud.BaseBarony
}

var BHassibenmadjeb哈西本马贾 feud.Barony = &哈西本马贾HassibenmadjebBarony{}

func init() {
    f := BHassibenmadjeb哈西本马贾.(*哈西本马贾HassibenmadjebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hassibenmadjeb",
		TitleName: "哈西本马贾",
		TitleCode: "b_hassibenmadjeb",
	}
}
