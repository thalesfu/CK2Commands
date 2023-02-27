package bohemia

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/bohemia/domazlice"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/bohemia/hradec"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/bohemia/litomerice"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/bohemia/plzen"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/bohemia/praha"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/bohemia/zatec"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BohemiaDuke interface {
    feud.Duke
    CDomazlice南波希米亚() 	domazlice.DomazliceCounty
    CHradec赫拉德茨() 	hradec.HradecCounty
    CLitomerice利托梅日采() 	litomerice.LitomericeCounty
    CPlzen普尔曾() 	plzen.PlzenCounty
    CPraha布拉格() 	praha.PrahaCounty
    CZatec扎泰茨() 	zatec.ZatecCounty
}

type 波希米亚BohemiaDuke struct {
	feud.BaseDuke
	南波希米亚Domazlice 	domazlice.DomazliceCounty
	赫拉德茨Hradec 	hradec.HradecCounty
	利托梅日采Litomerice 	litomerice.LitomericeCounty
	普尔曾Plzen 	plzen.PlzenCounty
	布拉格Praha 	praha.PrahaCounty
	扎泰茨Zatec 	zatec.ZatecCounty
}

func (d *波希米亚BohemiaDuke) CDomazlice南波希米亚() domazlice.DomazliceCounty {
	return d.南波希米亚Domazlice
}
    
func (d *波希米亚BohemiaDuke) CHradec赫拉德茨() hradec.HradecCounty {
	return d.赫拉德茨Hradec
}
    
func (d *波希米亚BohemiaDuke) CLitomerice利托梅日采() litomerice.LitomericeCounty {
	return d.利托梅日采Litomerice
}
    
func (d *波希米亚BohemiaDuke) CPlzen普尔曾() plzen.PlzenCounty {
	return d.普尔曾Plzen
}
    
func (d *波希米亚BohemiaDuke) CPraha布拉格() praha.PrahaCounty {
	return d.布拉格Praha
}
    
func (d *波希米亚BohemiaDuke) CZatec扎泰茨() zatec.ZatecCounty {
	return d.扎泰茨Zatec
}
    
var DBohemia波希米亚 BohemiaDuke = &波希米亚BohemiaDuke{}

func init() {
	f := DBohemia波希米亚.(*波希米亚BohemiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bohemia",
		TitleName: "波希米亚",
		TitleCode: "d_bohemia",
		Counties:  map[string]feud.County{},
	}

	f.南波希米亚Domazlice = domazlice.CDomazlice南波希米亚
	f.南波希米亚Domazlice.SetParent(f)
	
	f.赫拉德茨Hradec = hradec.CHradec赫拉德茨
	f.赫拉德茨Hradec.SetParent(f)
	
	f.利托梅日采Litomerice = litomerice.CLitomerice利托梅日采
	f.利托梅日采Litomerice.SetParent(f)
	
	f.普尔曾Plzen = plzen.CPlzen普尔曾
	f.普尔曾Plzen.SetParent(f)
	
	f.布拉格Praha = praha.CPraha布拉格
	f.布拉格Praha.SetParent(f)
	
	f.扎泰茨Zatec = zatec.CZatec扎泰茨
	f.扎泰茨Zatec.SetParent(f)
	
}
