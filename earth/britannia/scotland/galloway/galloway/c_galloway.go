package galloway

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GallowayCounty interface {
    feud.County
    BDumfries邓弗里斯() 	feud.Barony
    BDunragit邓拉吉特() 	feud.Barony
    BDunrod邓罗德() 	feud.Barony
    BGlenluce格伦卢斯() 	feud.Barony
    BKirkcudbright柯库布里() 	feud.Barony
    BThreave特利维() 	feud.Barony
    BWhithorn惠特霍恩() 	feud.Barony
    BWigtown威格敦() 	feud.Barony
}

type 加洛韦GallowayCounty struct {
	feud.BaseCounty
	邓弗里斯Dumfries 	feud.Barony
	邓拉吉特Dunragit 	feud.Barony
	邓罗德Dunrod 	feud.Barony
	格伦卢斯Glenluce 	feud.Barony
	柯库布里Kirkcudbright 	feud.Barony
	特利维Threave 	feud.Barony
	惠特霍恩Whithorn 	feud.Barony
	威格敦Wigtown 	feud.Barony
}

func (c *加洛韦GallowayCounty) BDumfries邓弗里斯() feud.Barony {
	return c.邓弗里斯Dumfries
}
    
func (c *加洛韦GallowayCounty) BDunragit邓拉吉特() feud.Barony {
	return c.邓拉吉特Dunragit
}
    
func (c *加洛韦GallowayCounty) BDunrod邓罗德() feud.Barony {
	return c.邓罗德Dunrod
}
    
func (c *加洛韦GallowayCounty) BGlenluce格伦卢斯() feud.Barony {
	return c.格伦卢斯Glenluce
}
    
func (c *加洛韦GallowayCounty) BKirkcudbright柯库布里() feud.Barony {
	return c.柯库布里Kirkcudbright
}
    
func (c *加洛韦GallowayCounty) BThreave特利维() feud.Barony {
	return c.特利维Threave
}
    
func (c *加洛韦GallowayCounty) BWhithorn惠特霍恩() feud.Barony {
	return c.惠特霍恩Whithorn
}
    
func (c *加洛韦GallowayCounty) BWigtown威格敦() feud.Barony {
	return c.威格敦Wigtown
}
    
var CGalloway加洛韦 GallowayCounty = &加洛韦GallowayCounty{}

func init() {
	f := CGalloway加洛韦.(*加洛韦GallowayCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "50",
		Title:     "galloway",
		TitleName: "加洛韦",
		TitleCode: "c_galloway",
		Baronies:  map[string]feud.Barony{},
	}

	f.邓弗里斯Dumfries = BDumfries邓弗里斯
	f.邓弗里斯Dumfries.SetParent(f)
	
	f.邓拉吉特Dunragit = BDunragit邓拉吉特
	f.邓拉吉特Dunragit.SetParent(f)
	
	f.邓罗德Dunrod = BDunrod邓罗德
	f.邓罗德Dunrod.SetParent(f)
	
	f.格伦卢斯Glenluce = BGlenluce格伦卢斯
	f.格伦卢斯Glenluce.SetParent(f)
	
	f.柯库布里Kirkcudbright = BKirkcudbright柯库布里
	f.柯库布里Kirkcudbright.SetParent(f)
	
	f.特利维Threave = BThreave特利维
	f.特利维Threave.SetParent(f)
	
	f.惠特霍恩Whithorn = BWhithorn惠特霍恩
	f.惠特霍恩Whithorn.SetParent(f)
	
	f.威格敦Wigtown = BWigtown威格敦
	f.威格敦Wigtown.SetParent(f)
	
}
