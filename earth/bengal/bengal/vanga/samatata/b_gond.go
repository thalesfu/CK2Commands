package samatata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡德GondBarony struct {
	feud.BaseBarony
}

var BGond贡德 feud.Barony = &贡德GondBarony{}

func init() {
	f := BGond贡德.(*贡德GondBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gond",
		TitleName: "贡德",
		TitleCode: "b_gond",
	}
}
