package shiraz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃格利德AzargartaBarony struct {
	feud.BaseBarony
}

var BAzargarta埃格利德 feud.Barony = &埃格利德AzargartaBarony{}

func init() {
    f := BAzargarta埃格利德.(*埃格利德AzargartaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azargarta",
		TitleName: "埃格利德",
		TitleCode: "b_azargarta",
	}
}
