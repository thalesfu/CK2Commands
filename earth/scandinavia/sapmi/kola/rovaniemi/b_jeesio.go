package rovaniemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶西厄JeesioBarony struct {
	feud.BaseBarony
}

var BJeesio耶西厄 feud.Barony = &耶西厄JeesioBarony{}

func init() {
    f := BJeesio耶西厄.(*耶西厄JeesioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jeesio",
		TitleName: "耶西厄",
		TitleCode: "b_jeesio",
	}
}
