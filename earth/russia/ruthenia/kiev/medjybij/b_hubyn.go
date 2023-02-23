package medjybij

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古宾HubynBarony struct {
	feud.BaseBarony
}

var BHubyn古宾 feud.Barony = &古宾HubynBarony{}

func init() {
	f := BHubyn古宾.(*古宾HubynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hubyn",
		TitleName: "古宾",
		TitleCode: "b_hubyn",
	}
}
