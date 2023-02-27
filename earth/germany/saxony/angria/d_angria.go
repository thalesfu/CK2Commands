package angria

import (
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/angria/minden"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/angria/oldenburg"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/angria/osnabruck"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AngriaDuke interface {
    feud.Duke
    CMinden明登() 	minden.MindenCounty
    COldenburg奥尔登堡() 	oldenburg.OldenburgCounty
    COsnabruck奥斯纳布吕克() 	osnabruck.OsnabruckCounty
}

type 盎格利亚AngriaDuke struct {
	feud.BaseDuke
	明登Minden 	minden.MindenCounty
	奥尔登堡Oldenburg 	oldenburg.OldenburgCounty
	奥斯纳布吕克Osnabruck 	osnabruck.OsnabruckCounty
}

func (d *盎格利亚AngriaDuke) CMinden明登() minden.MindenCounty {
	return d.明登Minden
}
    
func (d *盎格利亚AngriaDuke) COldenburg奥尔登堡() oldenburg.OldenburgCounty {
	return d.奥尔登堡Oldenburg
}
    
func (d *盎格利亚AngriaDuke) COsnabruck奥斯纳布吕克() osnabruck.OsnabruckCounty {
	return d.奥斯纳布吕克Osnabruck
}
    
var DAngria盎格利亚 AngriaDuke = &盎格利亚AngriaDuke{}

func init() {
	f := DAngria盎格利亚.(*盎格利亚AngriaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "angria",
		TitleName: "盎格利亚",
		TitleCode: "d_angria",
		Counties:  map[string]feud.County{},
	}

	f.明登Minden = minden.CMinden明登
	f.明登Minden.SetParent(f)
	
	f.奥尔登堡Oldenburg = oldenburg.COldenburg奥尔登堡
	f.奥尔登堡Oldenburg.SetParent(f)
	
	f.奥斯纳布吕克Osnabruck = osnabruck.COsnabruck奥斯纳布吕克
	f.奥斯纳布吕克Osnabruck.SetParent(f)
	
}
