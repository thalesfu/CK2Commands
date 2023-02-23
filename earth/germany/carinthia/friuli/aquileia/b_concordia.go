package aquileia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康科迪亚ConcordiaBarony struct {
	feud.BaseBarony
}

var BConcordia康科迪亚 feud.Barony = &康科迪亚ConcordiaBarony{}

func init() {
	f := BConcordia康科迪亚.(*康科迪亚ConcordiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "concordia",
		TitleName: "康科迪亚",
		TitleCode: "b_concordia",
	}
}
