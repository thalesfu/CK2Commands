package vanga

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/vanga/bikrampur"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/vanga/candradvipa"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/vanga/karmanta"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/vanga/samatata"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VangaDuke interface {
	feud.Duke
	CBikrampur毗讫罗摩补罗() bikrampur.BikrampurCounty
	CCandradvipa月洲() candradvipa.CandradvipaCounty
	CKarmanta羯曼多() karmanta.KarmantaCounty
	CSamatata三摩呾吒() samatata.SamatataCounty
}

type 旁伽VangaDuke struct {
	feud.BaseDuke
	毗讫罗摩补罗Bikrampur bikrampur.BikrampurCounty
	月洲Candradvipa   candradvipa.CandradvipaCounty
	羯曼多Karmanta     karmanta.KarmantaCounty
	三摩呾吒Samatata    samatata.SamatataCounty
}

func (d *旁伽VangaDuke) CBikrampur毗讫罗摩补罗() bikrampur.BikrampurCounty {
	return d.毗讫罗摩补罗Bikrampur
}

func (d *旁伽VangaDuke) CCandradvipa月洲() candradvipa.CandradvipaCounty {
	return d.月洲Candradvipa
}

func (d *旁伽VangaDuke) CKarmanta羯曼多() karmanta.KarmantaCounty {
	return d.羯曼多Karmanta
}

func (d *旁伽VangaDuke) CSamatata三摩呾吒() samatata.SamatataCounty {
	return d.三摩呾吒Samatata
}

var DVanga旁伽 VangaDuke = &旁伽VangaDuke{}

func init() {
	f := DVanga旁伽.(*旁伽VangaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "vanga",
		TitleName: "旁伽",
		TitleCode: "d_vanga",
		Counties:  map[string]feud.County{},
	}

	f.毗讫罗摩补罗Bikrampur = bikrampur.CBikrampur毗讫罗摩补罗
	f.毗讫罗摩补罗Bikrampur.SetParent(f)

	f.月洲Candradvipa = candradvipa.CCandradvipa月洲
	f.月洲Candradvipa.SetParent(f)

	f.羯曼多Karmanta = karmanta.CKarmanta羯曼多
	f.羯曼多Karmanta.SetParent(f)

	f.三摩呾吒Samatata = samatata.CSamatata三摩呾吒
	f.三摩呾吒Samatata.SetParent(f)

}
