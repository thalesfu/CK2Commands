package nilagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿努拉贾达AvulagaddaBarony struct {
	feud.BaseBarony
}

var BAvulagadda阿努拉贾达 feud.Barony = &阿努拉贾达AvulagaddaBarony{}

func init() {
	f := BAvulagadda阿努拉贾达.(*阿努拉贾达AvulagaddaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avulagadda",
		TitleName: "阿努拉贾达",
		TitleCode: "b_avulagadda",
	}
}
