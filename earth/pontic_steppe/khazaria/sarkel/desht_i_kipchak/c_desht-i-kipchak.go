package desht_i_kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Desht_i_kipchakCounty interface {
    feud.County
    BBakhmut巴赫穆特() 	feud.Barony
    BDobropillia多布罗皮利亚() 	feud.Barony
    BDruzhkivka德鲁日基夫卡() 	feud.Barony
    BKramatorsk克拉马托尔斯克() 	feud.Barony
    BKrasne克拉斯内() 	feud.Barony
    BLyman利曼() 	feud.Barony
    BMospyne莫斯皮诺() 	feud.Barony
    BSviatohirsk斯维亚托戈尔斯克() 	feud.Barony
}

type 钦察草原Desht_i_kipchakCounty struct {
	feud.BaseCounty
	巴赫穆特Bakhmut 	feud.Barony
	多布罗皮利亚Dobropillia 	feud.Barony
	德鲁日基夫卡Druzhkivka 	feud.Barony
	克拉马托尔斯克Kramatorsk 	feud.Barony
	克拉斯内Krasne 	feud.Barony
	利曼Lyman 	feud.Barony
	莫斯皮诺Mospyne 	feud.Barony
	斯维亚托戈尔斯克Sviatohirsk 	feud.Barony
}

func (c *钦察草原Desht_i_kipchakCounty) BBakhmut巴赫穆特() feud.Barony {
	return c.巴赫穆特Bakhmut
}
    
func (c *钦察草原Desht_i_kipchakCounty) BDobropillia多布罗皮利亚() feud.Barony {
	return c.多布罗皮利亚Dobropillia
}
    
func (c *钦察草原Desht_i_kipchakCounty) BDruzhkivka德鲁日基夫卡() feud.Barony {
	return c.德鲁日基夫卡Druzhkivka
}
    
func (c *钦察草原Desht_i_kipchakCounty) BKramatorsk克拉马托尔斯克() feud.Barony {
	return c.克拉马托尔斯克Kramatorsk
}
    
func (c *钦察草原Desht_i_kipchakCounty) BKrasne克拉斯内() feud.Barony {
	return c.克拉斯内Krasne
}
    
func (c *钦察草原Desht_i_kipchakCounty) BLyman利曼() feud.Barony {
	return c.利曼Lyman
}
    
func (c *钦察草原Desht_i_kipchakCounty) BMospyne莫斯皮诺() feud.Barony {
	return c.莫斯皮诺Mospyne
}
    
func (c *钦察草原Desht_i_kipchakCounty) BSviatohirsk斯维亚托戈尔斯克() feud.Barony {
	return c.斯维亚托戈尔斯克Sviatohirsk
}
    
var CDesht_i_kipchak钦察草原 Desht_i_kipchakCounty = &钦察草原Desht_i_kipchakCounty{}

func init() {
	f := CDesht_i_kipchak钦察草原.(*钦察草原Desht_i_kipchakCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "564",
		Title:     "desht_i_kipchak",
		TitleName: "钦察草原",
		TitleCode: "c_desht-i-kipchak",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴赫穆特Bakhmut = BBakhmut巴赫穆特
	f.巴赫穆特Bakhmut.SetParent(f)
	
	f.多布罗皮利亚Dobropillia = BDobropillia多布罗皮利亚
	f.多布罗皮利亚Dobropillia.SetParent(f)
	
	f.德鲁日基夫卡Druzhkivka = BDruzhkivka德鲁日基夫卡
	f.德鲁日基夫卡Druzhkivka.SetParent(f)
	
	f.克拉马托尔斯克Kramatorsk = BKramatorsk克拉马托尔斯克
	f.克拉马托尔斯克Kramatorsk.SetParent(f)
	
	f.克拉斯内Krasne = BKrasne克拉斯内
	f.克拉斯内Krasne.SetParent(f)
	
	f.利曼Lyman = BLyman利曼
	f.利曼Lyman.SetParent(f)
	
	f.莫斯皮诺Mospyne = BMospyne莫斯皮诺
	f.莫斯皮诺Mospyne.SetParent(f)
	
	f.斯维亚托戈尔斯克Sviatohirsk = BSviatohirsk斯维亚托戈尔斯克
	f.斯维亚托戈尔斯克Sviatohirsk.SetParent(f)
	
}
