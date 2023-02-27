package madhupur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆迦婆利诃BaghabarihatBarony struct {
	feud.BaseBarony
}

var BBaghabarihat婆迦婆利诃 feud.Barony = &婆迦婆利诃BaghabarihatBarony{}

func init() {
    f := BBaghabarihat婆迦婆利诃.(*婆迦婆利诃BaghabarihatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baghabarihat",
		TitleName: "婆迦婆利诃",
		TitleCode: "b_baghabarihat",
	}
}
