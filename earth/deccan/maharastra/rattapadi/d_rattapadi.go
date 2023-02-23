package rattapadi

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/rattapadi/kolhapur"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/rattapadi/lattalura"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/rattapadi/naldurg"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/rattapadi/taradavadi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RattapadiDuke interface {
	feud.Duke
	CKolhapur拘罗诃补罗() kolhapur.KolhapurCounty
	CLattalura罗多楼罗() lattalura.LattaluraCounty
	CNaldurg那罗突伽() naldurg.NaldurgCounty
	CTaradavadi多罗陀婆稚() taradavadi.TaradavadiCounty
}

type 剌吒波地RattapadiDuke struct {
	feud.BaseDuke
	拘罗诃补罗Kolhapur   kolhapur.KolhapurCounty
	罗多楼罗Lattalura   lattalura.LattaluraCounty
	那罗突伽Naldurg     naldurg.NaldurgCounty
	多罗陀婆稚Taradavadi taradavadi.TaradavadiCounty
}

func (d *剌吒波地RattapadiDuke) CKolhapur拘罗诃补罗() kolhapur.KolhapurCounty {
	return d.拘罗诃补罗Kolhapur
}

func (d *剌吒波地RattapadiDuke) CLattalura罗多楼罗() lattalura.LattaluraCounty {
	return d.罗多楼罗Lattalura
}

func (d *剌吒波地RattapadiDuke) CNaldurg那罗突伽() naldurg.NaldurgCounty {
	return d.那罗突伽Naldurg
}

func (d *剌吒波地RattapadiDuke) CTaradavadi多罗陀婆稚() taradavadi.TaradavadiCounty {
	return d.多罗陀婆稚Taradavadi
}

var DRattapadi剌吒波地 RattapadiDuke = &剌吒波地RattapadiDuke{}

func init() {
	f := DRattapadi剌吒波地.(*剌吒波地RattapadiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "rattapadi",
		TitleName: "剌吒波地",
		TitleCode: "d_rattapadi",
		Counties:  map[string]feud.County{},
	}

	f.拘罗诃补罗Kolhapur = kolhapur.CKolhapur拘罗诃补罗
	f.拘罗诃补罗Kolhapur.SetParent(f)

	f.罗多楼罗Lattalura = lattalura.CLattalura罗多楼罗
	f.罗多楼罗Lattalura.SetParent(f)

	f.那罗突伽Naldurg = naldurg.CNaldurg那罗突伽
	f.那罗突伽Naldurg.SetParent(f)

	f.多罗陀婆稚Taradavadi = taradavadi.CTaradavadi多罗陀婆稚
	f.多罗陀婆稚Taradavadi.SetParent(f)

}
