package nulambavadi

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/nulambavadi/alampur"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/nulambavadi/dwarasamudra"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/nulambavadi/uchangidurga"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NulambavadiDuke interface {
	feud.Duke
	CAlampur阿蓝补罗() alampur.AlampurCounty
	CDwarasamudra堕罗三慕达罗() dwarasamudra.DwarasamudraCounty
	CUchangidurga郁昌耆突伽() uchangidurga.UchangidurgaCounty
}

type 奴蓝婆婆稚NulambavadiDuke struct {
	feud.BaseDuke
	阿蓝补罗Alampur        alampur.AlampurCounty
	堕罗三慕达罗Dwarasamudra dwarasamudra.DwarasamudraCounty
	郁昌耆突伽Uchangidurga  uchangidurga.UchangidurgaCounty
}

func (d *奴蓝婆婆稚NulambavadiDuke) CAlampur阿蓝补罗() alampur.AlampurCounty {
	return d.阿蓝补罗Alampur
}

func (d *奴蓝婆婆稚NulambavadiDuke) CDwarasamudra堕罗三慕达罗() dwarasamudra.DwarasamudraCounty {
	return d.堕罗三慕达罗Dwarasamudra
}

func (d *奴蓝婆婆稚NulambavadiDuke) CUchangidurga郁昌耆突伽() uchangidurga.UchangidurgaCounty {
	return d.郁昌耆突伽Uchangidurga
}

var DNulambavadi奴蓝婆婆稚 NulambavadiDuke = &奴蓝婆婆稚NulambavadiDuke{}

func init() {
	f := DNulambavadi奴蓝婆婆稚.(*奴蓝婆婆稚NulambavadiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nulambavadi",
		TitleName: "奴蓝婆婆稚",
		TitleCode: "d_nulambavadi",
		Counties:  map[string]feud.County{},
	}

	f.阿蓝补罗Alampur = alampur.CAlampur阿蓝补罗
	f.阿蓝补罗Alampur.SetParent(f)

	f.堕罗三慕达罗Dwarasamudra = dwarasamudra.CDwarasamudra堕罗三慕达罗
	f.堕罗三慕达罗Dwarasamudra.SetParent(f)

	f.郁昌耆突伽Uchangidurga = uchangidurga.CUchangidurga郁昌耆突伽
	f.郁昌耆突伽Uchangidurga.SetParent(f)

}
