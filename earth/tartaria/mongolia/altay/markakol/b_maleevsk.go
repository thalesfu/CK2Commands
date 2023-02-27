package markakol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马列耶夫斯克MaleevskBarony struct {
	feud.BaseBarony
}

var BMaleevsk马列耶夫斯克 feud.Barony = &马列耶夫斯克MaleevskBarony{}

func init() {
    f := BMaleevsk马列耶夫斯克.(*马列耶夫斯克MaleevskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maleevsk",
		TitleName: "马列耶夫斯克",
		TitleCode: "b_maleevsk",
	}
}
