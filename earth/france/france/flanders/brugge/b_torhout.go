package brugge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔豪特TorhoutBarony struct {
	feud.BaseBarony
}

var BTorhout托尔豪特 feud.Barony = &托尔豪特TorhoutBarony{}

func init() {
	f := BTorhout托尔豪特.(*托尔豪特TorhoutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torhout",
		TitleName: "托尔豪特",
		TitleCode: "b_torhout",
	}
}
