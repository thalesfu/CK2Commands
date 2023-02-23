package saurashtra

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/saurashtra/bhumilka"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/saurashtra/somnath"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/saurashtra/valabhi"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/saurashtra/vardhamana"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaurashtraDuke interface {
	feud.Duke
	CBhumilka菩弥伽() bhumilka.BhumilkaCounty
	CSomnath须门那() somnath.SomnathCounty
	CValabhi伐腊毗() valabhi.ValabhiCounty
	CVardhamana筏陀摩那() vardhamana.VardhamanaCounty
}

type 苏剌侘SaurashtraDuke struct {
	feud.BaseDuke
	菩弥伽Bhumilka    bhumilka.BhumilkaCounty
	须门那Somnath     somnath.SomnathCounty
	伐腊毗Valabhi     valabhi.ValabhiCounty
	筏陀摩那Vardhamana vardhamana.VardhamanaCounty
}

func (d *苏剌侘SaurashtraDuke) CBhumilka菩弥伽() bhumilka.BhumilkaCounty {
	return d.菩弥伽Bhumilka
}

func (d *苏剌侘SaurashtraDuke) CSomnath须门那() somnath.SomnathCounty {
	return d.须门那Somnath
}

func (d *苏剌侘SaurashtraDuke) CValabhi伐腊毗() valabhi.ValabhiCounty {
	return d.伐腊毗Valabhi
}

func (d *苏剌侘SaurashtraDuke) CVardhamana筏陀摩那() vardhamana.VardhamanaCounty {
	return d.筏陀摩那Vardhamana
}

var DSaurashtra苏剌侘 SaurashtraDuke = &苏剌侘SaurashtraDuke{}

func init() {
	f := DSaurashtra苏剌侘.(*苏剌侘SaurashtraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "saurashtra",
		TitleName: "苏剌侘",
		TitleCode: "d_saurashtra",
		Counties:  map[string]feud.County{},
	}

	f.菩弥伽Bhumilka = bhumilka.CBhumilka菩弥伽
	f.菩弥伽Bhumilka.SetParent(f)

	f.须门那Somnath = somnath.CSomnath须门那
	f.须门那Somnath.SetParent(f)

	f.伐腊毗Valabhi = valabhi.CValabhi伐腊毗
	f.伐腊毗Valabhi.SetParent(f)

	f.筏陀摩那Vardhamana = vardhamana.CVardhamana筏陀摩那
	f.筏陀摩那Vardhamana.SetParent(f)

}
