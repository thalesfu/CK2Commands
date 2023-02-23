package hampshire

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/hampshire/dorset"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/hampshire/wight"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/hampshire/winchester"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HampshireDuke interface {
	feud.Duke
	CDorset多塞特() dorset.DorsetCounty
	CWight怀特() wight.WightCounty
	CWinchester温彻斯特() winchester.WinchesterCounty
}

type 威塞克斯HampshireDuke struct {
	feud.BaseDuke
	多塞特Dorset      dorset.DorsetCounty
	怀特Wight        wight.WightCounty
	温彻斯特Winchester winchester.WinchesterCounty
}

func (d *威塞克斯HampshireDuke) CDorset多塞特() dorset.DorsetCounty {
	return d.多塞特Dorset
}

func (d *威塞克斯HampshireDuke) CWight怀特() wight.WightCounty {
	return d.怀特Wight
}

func (d *威塞克斯HampshireDuke) CWinchester温彻斯特() winchester.WinchesterCounty {
	return d.温彻斯特Winchester
}

var DHampshire威塞克斯 HampshireDuke = &威塞克斯HampshireDuke{}

func init() {
	f := DHampshire威塞克斯.(*威塞克斯HampshireDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "hampshire",
		TitleName: "威塞克斯",
		TitleCode: "d_hampshire",
		Counties:  map[string]feud.County{},
	}

	f.多塞特Dorset = dorset.CDorset多塞特
	f.多塞特Dorset.SetParent(f)

	f.怀特Wight = wight.CWight怀特
	f.怀特Wight.SetParent(f)

	f.温彻斯特Winchester = winchester.CWinchester温彻斯特
	f.温彻斯特Winchester.SetParent(f)

}
