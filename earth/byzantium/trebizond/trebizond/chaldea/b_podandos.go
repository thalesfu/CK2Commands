package chaldea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波丹多斯PodandosBarony struct {
	feud.BaseBarony
}

var BPodandos波丹多斯 feud.Barony = &波丹多斯PodandosBarony{}

func init() {
    f := BPodandos波丹多斯.(*波丹多斯PodandosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "podandos",
		TitleName: "波丹多斯",
		TitleCode: "b_podandos",
	}
}
