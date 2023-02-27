package gurma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 滕科多戈TenkudugoBarony struct {
	feud.BaseBarony
}

var BTenkudugo滕科多戈 feud.Barony = &滕科多戈TenkudugoBarony{}

func init() {
    f := BTenkudugo滕科多戈.(*滕科多戈TenkudugoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tenkudugo",
		TitleName: "滕科多戈",
		TitleCode: "b_tenkudugo",
	}
}
