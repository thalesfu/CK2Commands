package qamdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QamdoCounty interface {
    feud.County
    BAsangka加桑卡() 	feud.Barony
    BGuro俄洛() 	feud.Barony
    BKarma嘎玛() 	feud.Barony
    BKarub卡若() 	feud.Barony
    BQamdo察木多() 	feud.Barony
    BRetong日通() 	feud.Barony
    BRiwoche类乌齐() 	feud.Barony
}

type 察木多QamdoCounty struct {
	feud.BaseCounty
	加桑卡Asangka 	feud.Barony
	俄洛Guro 	feud.Barony
	嘎玛Karma 	feud.Barony
	卡若Karub 	feud.Barony
	察木多Qamdo 	feud.Barony
	日通Retong 	feud.Barony
	类乌齐Riwoche 	feud.Barony
}

func (c *察木多QamdoCounty) BAsangka加桑卡() feud.Barony {
	return c.加桑卡Asangka
}
    
func (c *察木多QamdoCounty) BGuro俄洛() feud.Barony {
	return c.俄洛Guro
}
    
func (c *察木多QamdoCounty) BKarma嘎玛() feud.Barony {
	return c.嘎玛Karma
}
    
func (c *察木多QamdoCounty) BKarub卡若() feud.Barony {
	return c.卡若Karub
}
    
func (c *察木多QamdoCounty) BQamdo察木多() feud.Barony {
	return c.察木多Qamdo
}
    
func (c *察木多QamdoCounty) BRetong日通() feud.Barony {
	return c.日通Retong
}
    
func (c *察木多QamdoCounty) BRiwoche类乌齐() feud.Barony {
	return c.类乌齐Riwoche
}
    
var CQamdo察木多 QamdoCounty = &察木多QamdoCounty{}

func init() {
	f := CQamdo察木多.(*察木多QamdoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1508",
		Title:     "qamdo",
		TitleName: "察木多",
		TitleCode: "c_qamdo",
		Baronies:  map[string]feud.Barony{},
	}

	f.加桑卡Asangka = BAsangka加桑卡
	f.加桑卡Asangka.SetParent(f)
	
	f.俄洛Guro = BGuro俄洛
	f.俄洛Guro.SetParent(f)
	
	f.嘎玛Karma = BKarma嘎玛
	f.嘎玛Karma.SetParent(f)
	
	f.卡若Karub = BKarub卡若
	f.卡若Karub.SetParent(f)
	
	f.察木多Qamdo = BQamdo察木多
	f.察木多Qamdo.SetParent(f)
	
	f.日通Retong = BRetong日通
	f.日通Retong.SetParent(f)
	
	f.类乌齐Riwoche = BRiwoche类乌齐
	f.类乌齐Riwoche.SetParent(f)
	
}
