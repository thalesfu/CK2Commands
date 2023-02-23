package tarragona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎布里尔斯CambrilsBarony struct {
	feud.BaseBarony
}

var BCambrils坎布里尔斯 feud.Barony = &坎布里尔斯CambrilsBarony{}

func init() {
	f := BCambrils坎布里尔斯.(*坎布里尔斯CambrilsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cambrils",
		TitleName: "坎布里尔斯",
		TitleCode: "b_cambrils",
	}
}
