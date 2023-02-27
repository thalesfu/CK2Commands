package kuban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波德库莫克PodkumokBarony struct {
	feud.BaseBarony
}

var BPodkumok波德库莫克 feud.Barony = &波德库莫克PodkumokBarony{}

func init() {
    f := BPodkumok波德库莫克.(*波德库莫克PodkumokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "podkumok",
		TitleName: "波德库莫克",
		TitleCode: "b_podkumok",
	}
}
