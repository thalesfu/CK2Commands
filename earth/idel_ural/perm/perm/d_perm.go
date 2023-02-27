package perm

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/perm/kolva"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/perm/komi"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/perm/kudymkar"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/perm/perm"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/perm/tynea"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/perm/udmurts"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PermDuke interface {
    feud.Duke
    CKolva科尔瓦() 	kolva.KolvaCounty
    CKomi皮利瓦() 	komi.KomiCounty
    CKudymkar库德姆卡尔() 	kudymkar.KudymkarCounty
    CPerm彼尔姆() 	perm.PermCounty
    CTynea特涅阿() 	tynea.TyneaCounty
    CUdmurts乌德穆尔特() 	udmurts.UdmurtsCounty
}

type 彼尔姆PermDuke struct {
	feud.BaseDuke
	科尔瓦Kolva 	kolva.KolvaCounty
	皮利瓦Komi 	komi.KomiCounty
	库德姆卡尔Kudymkar 	kudymkar.KudymkarCounty
	彼尔姆Perm 	perm.PermCounty
	特涅阿Tynea 	tynea.TyneaCounty
	乌德穆尔特Udmurts 	udmurts.UdmurtsCounty
}

func (d *彼尔姆PermDuke) CKolva科尔瓦() kolva.KolvaCounty {
	return d.科尔瓦Kolva
}
    
func (d *彼尔姆PermDuke) CKomi皮利瓦() komi.KomiCounty {
	return d.皮利瓦Komi
}
    
func (d *彼尔姆PermDuke) CKudymkar库德姆卡尔() kudymkar.KudymkarCounty {
	return d.库德姆卡尔Kudymkar
}
    
func (d *彼尔姆PermDuke) CPerm彼尔姆() perm.PermCounty {
	return d.彼尔姆Perm
}
    
func (d *彼尔姆PermDuke) CTynea特涅阿() tynea.TyneaCounty {
	return d.特涅阿Tynea
}
    
func (d *彼尔姆PermDuke) CUdmurts乌德穆尔特() udmurts.UdmurtsCounty {
	return d.乌德穆尔特Udmurts
}
    
var DPerm彼尔姆 PermDuke = &彼尔姆PermDuke{}

func init() {
	f := DPerm彼尔姆.(*彼尔姆PermDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "perm",
		TitleName: "彼尔姆",
		TitleCode: "d_perm",
		Counties:  map[string]feud.County{},
	}

	f.科尔瓦Kolva = kolva.CKolva科尔瓦
	f.科尔瓦Kolva.SetParent(f)
	
	f.皮利瓦Komi = komi.CKomi皮利瓦
	f.皮利瓦Komi.SetParent(f)
	
	f.库德姆卡尔Kudymkar = kudymkar.CKudymkar库德姆卡尔
	f.库德姆卡尔Kudymkar.SetParent(f)
	
	f.彼尔姆Perm = perm.CPerm彼尔姆
	f.彼尔姆Perm.SetParent(f)
	
	f.特涅阿Tynea = tynea.CTynea特涅阿
	f.特涅阿Tynea.SetParent(f)
	
	f.乌德穆尔特Udmurts = udmurts.CUdmurts乌德穆尔特
	f.乌德穆尔特Udmurts.SetParent(f)
	
}
