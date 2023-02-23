package kermanshah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马赫库发MahalkufaBarony struct {
	feud.BaseBarony
}

var BMahalkufa马赫库发 feud.Barony = &马赫库发MahalkufaBarony{}

func init() {
	f := BMahalkufa马赫库发.(*马赫库发MahalkufaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahalkufa",
		TitleName: "马赫库发",
		TitleCode: "b_mahalkufa",
	}
}
