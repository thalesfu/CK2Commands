package qom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QomCounty interface {
    feud.County
    BDastjerd达斯特杰尔德() 	feud.Barony
    BJafariyeh贾法里耶赫() 	feud.Barony
    BJamkaran扎卡兰() 	feud.Barony
    BKahak卡哈克() 	feud.Barony
    BKhourabad豪尔萨巴德() 	feud.Barony
    BQanavat喀纳瓦特() 	feud.Barony
    BQom库姆() 	feud.Barony
    BSalafchegan萨拉彻甘() 	feud.Barony
}

type 库姆QomCounty struct {
	feud.BaseCounty
	达斯特杰尔德Dastjerd 	feud.Barony
	贾法里耶赫Jafariyeh 	feud.Barony
	扎卡兰Jamkaran 	feud.Barony
	卡哈克Kahak 	feud.Barony
	豪尔萨巴德Khourabad 	feud.Barony
	喀纳瓦特Qanavat 	feud.Barony
	库姆Qom 	feud.Barony
	萨拉彻甘Salafchegan 	feud.Barony
}

func (c *库姆QomCounty) BDastjerd达斯特杰尔德() feud.Barony {
	return c.达斯特杰尔德Dastjerd
}
    
func (c *库姆QomCounty) BJafariyeh贾法里耶赫() feud.Barony {
	return c.贾法里耶赫Jafariyeh
}
    
func (c *库姆QomCounty) BJamkaran扎卡兰() feud.Barony {
	return c.扎卡兰Jamkaran
}
    
func (c *库姆QomCounty) BKahak卡哈克() feud.Barony {
	return c.卡哈克Kahak
}
    
func (c *库姆QomCounty) BKhourabad豪尔萨巴德() feud.Barony {
	return c.豪尔萨巴德Khourabad
}
    
func (c *库姆QomCounty) BQanavat喀纳瓦特() feud.Barony {
	return c.喀纳瓦特Qanavat
}
    
func (c *库姆QomCounty) BQom库姆() feud.Barony {
	return c.库姆Qom
}
    
func (c *库姆QomCounty) BSalafchegan萨拉彻甘() feud.Barony {
	return c.萨拉彻甘Salafchegan
}
    
var CQom库姆 QomCounty = &库姆QomCounty{}

func init() {
	f := CQom库姆.(*库姆QomCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "659",
		Title:     "qom",
		TitleName: "库姆",
		TitleCode: "c_qom",
		Baronies:  map[string]feud.Barony{},
	}

	f.达斯特杰尔德Dastjerd = BDastjerd达斯特杰尔德
	f.达斯特杰尔德Dastjerd.SetParent(f)
	
	f.贾法里耶赫Jafariyeh = BJafariyeh贾法里耶赫
	f.贾法里耶赫Jafariyeh.SetParent(f)
	
	f.扎卡兰Jamkaran = BJamkaran扎卡兰
	f.扎卡兰Jamkaran.SetParent(f)
	
	f.卡哈克Kahak = BKahak卡哈克
	f.卡哈克Kahak.SetParent(f)
	
	f.豪尔萨巴德Khourabad = BKhourabad豪尔萨巴德
	f.豪尔萨巴德Khourabad.SetParent(f)
	
	f.喀纳瓦特Qanavat = BQanavat喀纳瓦特
	f.喀纳瓦特Qanavat.SetParent(f)
	
	f.库姆Qom = BQom库姆
	f.库姆Qom.SetParent(f)
	
	f.萨拉彻甘Salafchegan = BSalafchegan萨拉彻甘
	f.萨拉彻甘Salafchegan.SetParent(f)
	
}
