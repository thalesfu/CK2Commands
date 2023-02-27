package mordva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔达托沃ArdatovoBarony struct {
	feud.BaseBarony
}

var BArdatovo阿尔达托沃 feud.Barony = &阿尔达托沃ArdatovoBarony{}

func init() {
    f := BArdatovo阿尔达托沃.(*阿尔达托沃ArdatovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ardatovo",
		TitleName: "阿尔达托沃",
		TitleCode: "b_ardatovo",
	}
}
