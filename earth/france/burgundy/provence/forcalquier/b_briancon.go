package forcalquier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里昂松BrianconBarony struct {
	feud.BaseBarony
}

var BBriancon布里昂松 feud.Barony = &布里昂松BrianconBarony{}

func init() {
    f := BBriancon布里昂松.(*布里昂松BrianconBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "briancon",
		TitleName: "布里昂松",
		TitleCode: "b_briancon",
	}
}
