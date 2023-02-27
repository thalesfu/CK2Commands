package wana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 周拿古JunaghurBarony struct {
	feud.BaseBarony
}

var BJunaghur周拿古 feud.Barony = &周拿古JunaghurBarony{}

func init() {
    f := BJunaghur周拿古.(*周拿古JunaghurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "junaghur",
		TitleName: "周拿古",
		TitleCode: "b_junaghur",
	}
}
