package qarazhyrya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔沙特Gulshat_qarazhyryaBarony struct {
	feud.BaseBarony
}

var BGulshat_qarazhyrya古尔沙特 feud.Barony = &古尔沙特Gulshat_qarazhyryaBarony{}

func init() {
    f := BGulshat_qarazhyrya古尔沙特.(*古尔沙特Gulshat_qarazhyryaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gulshat_qarazhyrya",
		TitleName: "古尔沙特",
		TitleCode: "b_gulshat_qarazhyrya",
	}
}
