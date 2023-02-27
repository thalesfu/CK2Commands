package zeta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波德戈里察PodgoricaBarony struct {
	feud.BaseBarony
}

var BPodgorica波德戈里察 feud.Barony = &波德戈里察PodgoricaBarony{}

func init() {
    f := BPodgorica波德戈里察.(*波德戈里察PodgoricaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "podgorica",
		TitleName: "波德戈里察",
		TitleCode: "b_podgorica",
	}
}
