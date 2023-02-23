package tadmor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古里亚AlqunjahBarony struct {
	feud.BaseBarony
}

var BAlqunjah古里亚 feud.Barony = &古里亚AlqunjahBarony{}

func init() {
	f := BAlqunjah古里亚.(*古里亚AlqunjahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alqunjah",
		TitleName: "古里亚",
		TitleCode: "b_alqunjah",
	}
}
