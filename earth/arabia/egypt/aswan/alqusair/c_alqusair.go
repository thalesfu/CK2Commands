package alqusair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlqusairCounty interface {
    feud.County
    BAbughusun古胡孙() 	feud.Barony
    BAsfoun阿斯丰() 	feud.Barony
    BBureid布赖德() 	feud.Barony
    BJamsah杰姆萨() 	feud.Barony
    BKosseir库赛尔() 	feud.Barony
    BSafaga塞法杰() 	feud.Barony
    BUmmrus乌姆鲁斯() 	feud.Barony
}

type 古赛尔AlqusairCounty struct {
	feud.BaseCounty
	古胡孙Abughusun 	feud.Barony
	阿斯丰Asfoun 	feud.Barony
	布赖德Bureid 	feud.Barony
	杰姆萨Jamsah 	feud.Barony
	库赛尔Kosseir 	feud.Barony
	塞法杰Safaga 	feud.Barony
	乌姆鲁斯Ummrus 	feud.Barony
}

func (c *古赛尔AlqusairCounty) BAbughusun古胡孙() feud.Barony {
	return c.古胡孙Abughusun
}
    
func (c *古赛尔AlqusairCounty) BAsfoun阿斯丰() feud.Barony {
	return c.阿斯丰Asfoun
}
    
func (c *古赛尔AlqusairCounty) BBureid布赖德() feud.Barony {
	return c.布赖德Bureid
}
    
func (c *古赛尔AlqusairCounty) BJamsah杰姆萨() feud.Barony {
	return c.杰姆萨Jamsah
}
    
func (c *古赛尔AlqusairCounty) BKosseir库赛尔() feud.Barony {
	return c.库赛尔Kosseir
}
    
func (c *古赛尔AlqusairCounty) BSafaga塞法杰() feud.Barony {
	return c.塞法杰Safaga
}
    
func (c *古赛尔AlqusairCounty) BUmmrus乌姆鲁斯() feud.Barony {
	return c.乌姆鲁斯Ummrus
}
    
var CAlqusair古赛尔 AlqusairCounty = &古赛尔AlqusairCounty{}

func init() {
	f := CAlqusair古赛尔.(*古赛尔AlqusairCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1781",
		Title:     "alqusair",
		TitleName: "古赛尔",
		TitleCode: "c_alqusair",
		Baronies:  map[string]feud.Barony{},
	}

	f.古胡孙Abughusun = BAbughusun古胡孙
	f.古胡孙Abughusun.SetParent(f)
	
	f.阿斯丰Asfoun = BAsfoun阿斯丰
	f.阿斯丰Asfoun.SetParent(f)
	
	f.布赖德Bureid = BBureid布赖德
	f.布赖德Bureid.SetParent(f)
	
	f.杰姆萨Jamsah = BJamsah杰姆萨
	f.杰姆萨Jamsah.SetParent(f)
	
	f.库赛尔Kosseir = BKosseir库赛尔
	f.库赛尔Kosseir.SetParent(f)
	
	f.塞法杰Safaga = BSafaga塞法杰
	f.塞法杰Safaga.SetParent(f)
	
	f.乌姆鲁斯Ummrus = BUmmrus乌姆鲁斯
	f.乌姆鲁斯Ummrus.SetParent(f)
	
}
