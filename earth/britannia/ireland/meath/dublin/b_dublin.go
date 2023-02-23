package dublin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 都柏林DublinBarony struct {
	feud.BaseBarony
}

var BDublin都柏林 feud.Barony = &都柏林DublinBarony{}

func init() {
	f := BDublin都柏林.(*都柏林DublinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dublin",
		TitleName: "都柏林",
		TitleCode: "b_dublin",
	}
}
