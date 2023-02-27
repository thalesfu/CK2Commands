package qarazhyrya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克套Aktau_qarazhyryaBarony struct {
	feud.BaseBarony
}

var BAktau_qarazhyrya阿克套 feud.Barony = &阿克套Aktau_qarazhyryaBarony{}

func init() {
    f := BAktau_qarazhyrya阿克套.(*阿克套Aktau_qarazhyryaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aktau_qarazhyrya",
		TitleName: "阿克套",
		TitleCode: "b_aktau_qarazhyrya",
	}
}
