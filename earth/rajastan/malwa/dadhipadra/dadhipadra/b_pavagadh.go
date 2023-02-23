package dadhipadra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波伐伽德PavagadhBarony struct {
	feud.BaseBarony
}

var BPavagadh波伐伽德 feud.Barony = &波伐伽德PavagadhBarony{}

func init() {
	f := BPavagadh波伐伽德.(*波伐伽德PavagadhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pavagadh",
		TitleName: "波伐伽德",
		TitleCode: "b_pavagadh",
	}
}
