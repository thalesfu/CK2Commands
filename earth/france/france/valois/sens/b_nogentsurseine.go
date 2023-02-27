package sens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞纳河畔诺让NogentsurseineBarony struct {
	feud.BaseBarony
}

var BNogentsurseine塞纳河畔诺让 feud.Barony = &塞纳河畔诺让NogentsurseineBarony{}

func init() {
    f := BNogentsurseine塞纳河畔诺让.(*塞纳河畔诺让NogentsurseineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nogentsurseine",
		TitleName: "塞纳河畔诺让",
		TitleCode: "b_nogentsurseine",
	}
}
