package dublin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克利Ath_cliathBarony struct {
	feud.BaseBarony
}

var BAth_cliath阿克利 feud.Barony = &阿克利Ath_cliathBarony{}

func init() {
    f := BAth_cliath阿克利.(*阿克利Ath_cliathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ath_cliath",
		TitleName: "阿克利",
		TitleCode: "b_ath_cliath",
	}
}
