package gnieznienskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波德加耶PodgajeBarony struct {
	feud.BaseBarony
}

var BPodgaje波德加耶 feud.Barony = &波德加耶PodgajeBarony{}

func init() {
    f := BPodgaje波德加耶.(*波德加耶PodgajeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "podgaje",
		TitleName: "波德加耶",
		TitleCode: "b_podgaje",
	}
}
