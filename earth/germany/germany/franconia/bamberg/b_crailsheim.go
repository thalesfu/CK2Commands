package bamberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克赖尔斯海姆CrailsheimBarony struct {
	feud.BaseBarony
}

var BCrailsheim克赖尔斯海姆 feud.Barony = &克赖尔斯海姆CrailsheimBarony{}

func init() {
    f := BCrailsheim克赖尔斯海姆.(*克赖尔斯海姆CrailsheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crailsheim",
		TitleName: "克赖尔斯海姆",
		TitleCode: "b_crailsheim",
	}
}
