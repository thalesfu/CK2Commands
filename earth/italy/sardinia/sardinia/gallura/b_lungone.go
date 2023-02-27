package gallura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伦贡LungoneBarony struct {
	feud.BaseBarony
}

var BLungone伦贡 feud.Barony = &伦贡LungoneBarony{}

func init() {
    f := BLungone伦贡.(*伦贡LungoneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lungone",
		TitleName: "伦贡",
		TitleCode: "b_lungone",
	}
}
