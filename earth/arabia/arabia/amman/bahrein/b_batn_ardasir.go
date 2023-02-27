package bahrein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴特恩阿尔达希尔Batn_ardasirBarony struct {
	feud.BaseBarony
}

var BBatn_ardasir巴特恩阿尔达希尔 feud.Barony = &巴特恩阿尔达希尔Batn_ardasirBarony{}

func init() {
    f := BBatn_ardasir巴特恩阿尔达希尔.(*巴特恩阿尔达希尔Batn_ardasirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "batn_ardasir",
		TitleName: "巴特恩阿尔达希尔",
		TitleCode: "b_batn_ardasir",
	}
}
