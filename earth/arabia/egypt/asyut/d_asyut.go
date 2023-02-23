package asyut

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/asyut/asyut"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/asyut/esna"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/asyut/kharga"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AsyutDuke interface {
	feud.Duke
	CAsyut艾斯尤特() asyut.AsyutCounty
	CEsna埃斯纳() esna.EsnaCounty
	CKharga哈里杰() kharga.KhargaCounty
}

type 艾斯尤特AsyutDuke struct {
	feud.BaseDuke
	艾斯尤特Asyut asyut.AsyutCounty
	埃斯纳Esna   esna.EsnaCounty
	哈里杰Kharga kharga.KhargaCounty
}

func (d *艾斯尤特AsyutDuke) CAsyut艾斯尤特() asyut.AsyutCounty {
	return d.艾斯尤特Asyut
}

func (d *艾斯尤特AsyutDuke) CEsna埃斯纳() esna.EsnaCounty {
	return d.埃斯纳Esna
}

func (d *艾斯尤特AsyutDuke) CKharga哈里杰() kharga.KhargaCounty {
	return d.哈里杰Kharga
}

var DAsyut艾斯尤特 AsyutDuke = &艾斯尤特AsyutDuke{}

func init() {
	f := DAsyut艾斯尤特.(*艾斯尤特AsyutDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "asyut",
		TitleName: "艾斯尤特",
		TitleCode: "d_asyut",
		Counties:  map[string]feud.County{},
	}

	f.艾斯尤特Asyut = asyut.CAsyut艾斯尤特
	f.艾斯尤特Asyut.SetParent(f)

	f.埃斯纳Esna = esna.CEsna埃斯纳
	f.埃斯纳Esna.SetParent(f)

	f.哈里杰Kharga = kharga.CKharga哈里杰
	f.哈里杰Kharga.SetParent(f)

}
