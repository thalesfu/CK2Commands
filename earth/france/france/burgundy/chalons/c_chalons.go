package chalons

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChalonsCounty interface {
    feud.County
    BBrancion布朗雄() 	feud.Barony
    BChalon沙隆() 	feud.Barony
    BChamilly沙米伊() 	feud.Barony
    BCuisery奎瑟里() 	feud.Barony
    BLouhans卢昂() 	feud.Barony
    BSeurre瑟尔() 	feud.Barony
    BStjeandelosne圣让德洛讷() 	feud.Barony
    BTournus图尔尼() 	feud.Barony
}

type 沙隆ChalonsCounty struct {
	feud.BaseCounty
	布朗雄Brancion 	feud.Barony
	沙隆Chalon 	feud.Barony
	沙米伊Chamilly 	feud.Barony
	奎瑟里Cuisery 	feud.Barony
	卢昂Louhans 	feud.Barony
	瑟尔Seurre 	feud.Barony
	圣让德洛讷Stjeandelosne 	feud.Barony
	图尔尼Tournus 	feud.Barony
}

func (c *沙隆ChalonsCounty) BBrancion布朗雄() feud.Barony {
	return c.布朗雄Brancion
}
    
func (c *沙隆ChalonsCounty) BChalon沙隆() feud.Barony {
	return c.沙隆Chalon
}
    
func (c *沙隆ChalonsCounty) BChamilly沙米伊() feud.Barony {
	return c.沙米伊Chamilly
}
    
func (c *沙隆ChalonsCounty) BCuisery奎瑟里() feud.Barony {
	return c.奎瑟里Cuisery
}
    
func (c *沙隆ChalonsCounty) BLouhans卢昂() feud.Barony {
	return c.卢昂Louhans
}
    
func (c *沙隆ChalonsCounty) BSeurre瑟尔() feud.Barony {
	return c.瑟尔Seurre
}
    
func (c *沙隆ChalonsCounty) BStjeandelosne圣让德洛讷() feud.Barony {
	return c.圣让德洛讷Stjeandelosne
}
    
func (c *沙隆ChalonsCounty) BTournus图尔尼() feud.Barony {
	return c.图尔尼Tournus
}
    
var CChalons沙隆 ChalonsCounty = &沙隆ChalonsCounty{}

func init() {
	f := CChalons沙隆.(*沙隆ChalonsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "240",
		Title:     "chalons",
		TitleName: "沙隆",
		TitleCode: "c_chalons",
		Baronies:  map[string]feud.Barony{},
	}

	f.布朗雄Brancion = BBrancion布朗雄
	f.布朗雄Brancion.SetParent(f)
	
	f.沙隆Chalon = BChalon沙隆
	f.沙隆Chalon.SetParent(f)
	
	f.沙米伊Chamilly = BChamilly沙米伊
	f.沙米伊Chamilly.SetParent(f)
	
	f.奎瑟里Cuisery = BCuisery奎瑟里
	f.奎瑟里Cuisery.SetParent(f)
	
	f.卢昂Louhans = BLouhans卢昂
	f.卢昂Louhans.SetParent(f)
	
	f.瑟尔Seurre = BSeurre瑟尔
	f.瑟尔Seurre.SetParent(f)
	
	f.圣让德洛讷Stjeandelosne = BStjeandelosne圣让德洛讷
	f.圣让德洛讷Stjeandelosne.SetParent(f)
	
	f.图尔尼Tournus = BTournus图尔尼
	f.图尔尼Tournus.SetParent(f)
	
}
