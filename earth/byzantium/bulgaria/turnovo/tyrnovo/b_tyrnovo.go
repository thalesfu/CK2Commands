package tyrnovo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔诺沃TyrnovoBarony struct {
	feud.BaseBarony
}

var BTyrnovo特尔诺沃 feud.Barony = &特尔诺沃TyrnovoBarony{}

func init() {
    f := BTyrnovo特尔诺沃.(*特尔诺沃TyrnovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyrnovo",
		TitleName: "特尔诺沃",
		TitleCode: "b_tyrnovo",
	}
}
