package badakhshan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BadakhshanCounty interface {
    feud.County
    BBadakhshan波多叉拏() 	feud.Barony
    BJarf贾尔夫() 	feud.Barony
    BJerm杰姆() 	feud.Barony
    BKamar卡马尔() 	feud.Barony
    BKerran凯朗() 	feud.Barony
    BPanj喷赤() 	feud.Barony
    BRustaq鲁斯塔克() 	feud.Barony
}

type 波多叉拏BadakhshanCounty struct {
	feud.BaseCounty
	波多叉拏Badakhshan 	feud.Barony
	贾尔夫Jarf 	feud.Barony
	杰姆Jerm 	feud.Barony
	卡马尔Kamar 	feud.Barony
	凯朗Kerran 	feud.Barony
	喷赤Panj 	feud.Barony
	鲁斯塔克Rustaq 	feud.Barony
}

func (c *波多叉拏BadakhshanCounty) BBadakhshan波多叉拏() feud.Barony {
	return c.波多叉拏Badakhshan
}
    
func (c *波多叉拏BadakhshanCounty) BJarf贾尔夫() feud.Barony {
	return c.贾尔夫Jarf
}
    
func (c *波多叉拏BadakhshanCounty) BJerm杰姆() feud.Barony {
	return c.杰姆Jerm
}
    
func (c *波多叉拏BadakhshanCounty) BKamar卡马尔() feud.Barony {
	return c.卡马尔Kamar
}
    
func (c *波多叉拏BadakhshanCounty) BKerran凯朗() feud.Barony {
	return c.凯朗Kerran
}
    
func (c *波多叉拏BadakhshanCounty) BPanj喷赤() feud.Barony {
	return c.喷赤Panj
}
    
func (c *波多叉拏BadakhshanCounty) BRustaq鲁斯塔克() feud.Barony {
	return c.鲁斯塔克Rustaq
}
    
var CBadakhshan波多叉拏 BadakhshanCounty = &波多叉拏BadakhshanCounty{}

func init() {
	f := CBadakhshan波多叉拏.(*波多叉拏BadakhshanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1547",
		Title:     "badakhshan",
		TitleName: "波多叉拏",
		TitleCode: "c_badakhshan",
		Baronies:  map[string]feud.Barony{},
	}

	f.波多叉拏Badakhshan = BBadakhshan波多叉拏
	f.波多叉拏Badakhshan.SetParent(f)
	
	f.贾尔夫Jarf = BJarf贾尔夫
	f.贾尔夫Jarf.SetParent(f)
	
	f.杰姆Jerm = BJerm杰姆
	f.杰姆Jerm.SetParent(f)
	
	f.卡马尔Kamar = BKamar卡马尔
	f.卡马尔Kamar.SetParent(f)
	
	f.凯朗Kerran = BKerran凯朗
	f.凯朗Kerran.SetParent(f)
	
	f.喷赤Panj = BPanj喷赤
	f.喷赤Panj.SetParent(f)
	
	f.鲁斯塔克Rustaq = BRustaq鲁斯塔克
	f.鲁斯塔克Rustaq.SetParent(f)
	
}
