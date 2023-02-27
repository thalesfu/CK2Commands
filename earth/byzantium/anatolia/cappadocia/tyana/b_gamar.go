package tyana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽马GamarBarony struct {
	feud.BaseBarony
}

var BGamar伽马 feud.Barony = &伽马GamarBarony{}

func init() {
    f := BGamar伽马.(*伽马GamarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gamar",
		TitleName: "伽马",
		TitleCode: "b_gamar",
	}
}
