package dakhina_desa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩喜扬伽那MahiyanganaBarony struct {
	feud.BaseBarony
}

var BMahiyangana摩喜扬伽那 feud.Barony = &摩喜扬伽那MahiyanganaBarony{}

func init() {
    f := BMahiyangana摩喜扬伽那.(*摩喜扬伽那MahiyanganaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahiyangana",
		TitleName: "摩喜扬伽那",
		TitleCode: "b_mahiyangana",
	}
}
