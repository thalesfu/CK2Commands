package ili

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亦力把里IlibaliqBarony struct {
	feud.BaseBarony
}

var BIlibaliq亦力把里 feud.Barony = &亦力把里IlibaliqBarony{}

func init() {
	f := BIlibaliq亦力把里.(*亦力把里IlibaliqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilibaliq",
		TitleName: "亦力把里",
		TitleCode: "b_ilibaliq",
	}
}
