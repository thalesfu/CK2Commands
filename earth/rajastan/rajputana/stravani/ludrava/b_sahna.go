package ludrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 犀诃那SahnaBarony struct {
	feud.BaseBarony
}

var BSahna犀诃那 feud.Barony = &犀诃那SahnaBarony{}

func init() {
    f := BSahna犀诃那.(*犀诃那SahnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sahna",
		TitleName: "犀诃那",
		TitleCode: "b_sahna",
	}
}
