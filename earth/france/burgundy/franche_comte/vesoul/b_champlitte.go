package vesoul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尚普利特ChamplitteBarony struct {
	feud.BaseBarony
}

var BChamplitte尚普利特 feud.Barony = &尚普利特ChamplitteBarony{}

func init() {
    f := BChamplitte尚普利特.(*尚普利特ChamplitteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "champlitte",
		TitleName: "尚普利特",
		TitleCode: "b_champlitte",
	}
}
