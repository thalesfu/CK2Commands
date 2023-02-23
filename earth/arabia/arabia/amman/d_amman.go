package amman

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/amman/bahrein"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/amman/damman"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/amman/kuwait"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/amman/uwal"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmmanDuke interface {
	feud.Duke
	CBahrein巴林() bahrein.BahreinCounty
	CDamman盖提夫() damman.DammanCounty
	CKuwait科威特() kuwait.KuwaitCounty
	CUwal乌瓦尔() uwal.UwalCounty
}

type 巴林AmmanDuke struct {
	feud.BaseDuke
	巴林Bahrein bahrein.BahreinCounty
	盖提夫Damman damman.DammanCounty
	科威特Kuwait kuwait.KuwaitCounty
	乌瓦尔Uwal   uwal.UwalCounty
}

func (d *巴林AmmanDuke) CBahrein巴林() bahrein.BahreinCounty {
	return d.巴林Bahrein
}

func (d *巴林AmmanDuke) CDamman盖提夫() damman.DammanCounty {
	return d.盖提夫Damman
}

func (d *巴林AmmanDuke) CKuwait科威特() kuwait.KuwaitCounty {
	return d.科威特Kuwait
}

func (d *巴林AmmanDuke) CUwal乌瓦尔() uwal.UwalCounty {
	return d.乌瓦尔Uwal
}

var DAmman巴林 AmmanDuke = &巴林AmmanDuke{}

func init() {
	f := DAmman巴林.(*巴林AmmanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "amman",
		TitleName: "巴林",
		TitleCode: "d_amman",
		Counties:  map[string]feud.County{},
	}

	f.巴林Bahrein = bahrein.CBahrein巴林
	f.巴林Bahrein.SetParent(f)

	f.盖提夫Damman = damman.CDamman盖提夫
	f.盖提夫Damman.SetParent(f)

	f.科威特Kuwait = kuwait.CKuwait科威特
	f.科威特Kuwait.SetParent(f)

	f.乌瓦尔Uwal = uwal.CUwal乌瓦尔
	f.乌瓦尔Uwal.SetParent(f)

}
