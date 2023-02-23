package breda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝亨奥普佐姆BergenopzoomBarony struct {
	feud.BaseBarony
}

var BBergenopzoom贝亨奥普佐姆 feud.Barony = &贝亨奥普佐姆BergenopzoomBarony{}

func init() {
	f := BBergenopzoom贝亨奥普佐姆.(*贝亨奥普佐姆BergenopzoomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bergenopzoom",
		TitleName: "贝亨奥普佐姆",
		TitleCode: "b_bergenopzoom",
	}
}
