package alcala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔戈_德桑蒂乌斯特Burgo_de_santiusteBarony struct {
	feud.BaseBarony
}

var BBurgo_de_santiuste布尔戈_德桑蒂乌斯特 feud.Barony = &布尔戈_德桑蒂乌斯特Burgo_de_santiusteBarony{}

func init() {
    f := BBurgo_de_santiuste布尔戈_德桑蒂乌斯特.(*布尔戈_德桑蒂乌斯特Burgo_de_santiusteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burgo_de_santiuste",
		TitleName: "布尔戈_德桑蒂乌斯特",
		TitleCode: "b_burgo_de_santiuste",
	}
}
