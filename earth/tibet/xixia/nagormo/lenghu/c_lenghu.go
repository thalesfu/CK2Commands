package lenghu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LenghuCounty interface {
    feud.County
    BJunheju君和居() 	feud.Barony
    BLedu乐都() 	feud.Barony
    BLenghu冷湖() 	feud.Barony
    BTuanjie团结() 	feud.Barony
    BXinghu星湖() 	feud.Barony
    BYanhua盐化() 	feud.Barony
    BYucai育才() 	feud.Barony
}

type 冷湖LenghuCounty struct {
	feud.BaseCounty
	君和居Junheju 	feud.Barony
	乐都Ledu 	feud.Barony
	冷湖Lenghu 	feud.Barony
	团结Tuanjie 	feud.Barony
	星湖Xinghu 	feud.Barony
	盐化Yanhua 	feud.Barony
	育才Yucai 	feud.Barony
}

func (c *冷湖LenghuCounty) BJunheju君和居() feud.Barony {
	return c.君和居Junheju
}
    
func (c *冷湖LenghuCounty) BLedu乐都() feud.Barony {
	return c.乐都Ledu
}
    
func (c *冷湖LenghuCounty) BLenghu冷湖() feud.Barony {
	return c.冷湖Lenghu
}
    
func (c *冷湖LenghuCounty) BTuanjie团结() feud.Barony {
	return c.团结Tuanjie
}
    
func (c *冷湖LenghuCounty) BXinghu星湖() feud.Barony {
	return c.星湖Xinghu
}
    
func (c *冷湖LenghuCounty) BYanhua盐化() feud.Barony {
	return c.盐化Yanhua
}
    
func (c *冷湖LenghuCounty) BYucai育才() feud.Barony {
	return c.育才Yucai
}
    
var CLenghu冷湖 LenghuCounty = &冷湖LenghuCounty{}

func init() {
	f := CLenghu冷湖.(*冷湖LenghuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1571",
		Title:     "lenghu",
		TitleName: "冷湖",
		TitleCode: "c_lenghu",
		Baronies:  map[string]feud.Barony{},
	}

	f.君和居Junheju = BJunheju君和居
	f.君和居Junheju.SetParent(f)
	
	f.乐都Ledu = BLedu乐都
	f.乐都Ledu.SetParent(f)
	
	f.冷湖Lenghu = BLenghu冷湖
	f.冷湖Lenghu.SetParent(f)
	
	f.团结Tuanjie = BTuanjie团结
	f.团结Tuanjie.SetParent(f)
	
	f.星湖Xinghu = BXinghu星湖
	f.星湖Xinghu.SetParent(f)
	
	f.盐化Yanhua = BYanhua盐化
	f.盐化Yanhua.SetParent(f)
	
	f.育才Yucai = BYucai育才
	f.育才Yucai.SetParent(f)
	
}
