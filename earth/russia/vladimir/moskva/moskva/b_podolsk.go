package moskva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波多利斯克PodolskBarony struct {
	feud.BaseBarony
}

var BPodolsk波多利斯克 feud.Barony = &波多利斯克PodolskBarony{}

func init() {
	f := BPodolsk波多利斯克.(*波多利斯克PodolskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "podolsk",
		TitleName: "波多利斯克",
		TitleCode: "b_podolsk",
	}
}
