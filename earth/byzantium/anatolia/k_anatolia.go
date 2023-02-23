package anatolia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/anatolia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/bucellarian"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/cappadocia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/charsianon"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/cibyrrhaeot"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/cyprus"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/nikaea"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/samos"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/thracesia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnatoliaKingdom interface {
	feud.Kingdom
	DAnatolia安纳托利亚() anatolia.AnatoliaDuke
	DBucellarian布凯拉里翁() bucellarian.BucellarianDuke
	DCappadocia卡帕多西亚() cappadocia.CappadociaDuke
	DCharsianon哈耳西亚农() charsianon.CharsianonDuke
	DCibyrrhaeot基比拉奥特() cibyrrhaeot.CibyrrhaeotDuke
	DCyprus塞浦路斯() cyprus.CyprusDuke
	DNikaea奥普西金() nikaea.NikaeaDuke
	DSamos萨摩斯() samos.SamosDuke
	DThracesia色雷斯西亚() thracesia.ThracesiaDuke
}

type 安纳托利亚AnatoliaKingdom struct {
	feud.BaseKingdom
	安纳托利亚Anatolia    anatolia.AnatoliaDuke
	布凯拉里翁Bucellarian bucellarian.BucellarianDuke
	卡帕多西亚Cappadocia  cappadocia.CappadociaDuke
	哈耳西亚农Charsianon  charsianon.CharsianonDuke
	基比拉奥特Cibyrrhaeot cibyrrhaeot.CibyrrhaeotDuke
	塞浦路斯Cyprus       cyprus.CyprusDuke
	奥普西金Nikaea       nikaea.NikaeaDuke
	萨摩斯Samos         samos.SamosDuke
	色雷斯西亚Thracesia   thracesia.ThracesiaDuke
}

func (k *安纳托利亚AnatoliaKingdom) DAnatolia安纳托利亚() anatolia.AnatoliaDuke {
	return k.安纳托利亚Anatolia
}

func (k *安纳托利亚AnatoliaKingdom) DBucellarian布凯拉里翁() bucellarian.BucellarianDuke {
	return k.布凯拉里翁Bucellarian
}

func (k *安纳托利亚AnatoliaKingdom) DCappadocia卡帕多西亚() cappadocia.CappadociaDuke {
	return k.卡帕多西亚Cappadocia
}

func (k *安纳托利亚AnatoliaKingdom) DCharsianon哈耳西亚农() charsianon.CharsianonDuke {
	return k.哈耳西亚农Charsianon
}

func (k *安纳托利亚AnatoliaKingdom) DCibyrrhaeot基比拉奥特() cibyrrhaeot.CibyrrhaeotDuke {
	return k.基比拉奥特Cibyrrhaeot
}

func (k *安纳托利亚AnatoliaKingdom) DCyprus塞浦路斯() cyprus.CyprusDuke {
	return k.塞浦路斯Cyprus
}

func (k *安纳托利亚AnatoliaKingdom) DNikaea奥普西金() nikaea.NikaeaDuke {
	return k.奥普西金Nikaea
}

func (k *安纳托利亚AnatoliaKingdom) DSamos萨摩斯() samos.SamosDuke {
	return k.萨摩斯Samos
}

func (k *安纳托利亚AnatoliaKingdom) DThracesia色雷斯西亚() thracesia.ThracesiaDuke {
	return k.色雷斯西亚Thracesia
}

var KAnatolia安纳托利亚 AnatoliaKingdom = &安纳托利亚AnatoliaKingdom{}

func init() {
	f := KAnatolia安纳托利亚.(*安纳托利亚AnatoliaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "anatolia",
		TitleName: "安纳托利亚",
		TitleCode: "k_anatolia",
		Dukes:     map[string]feud.Duke{},
	}

	f.安纳托利亚Anatolia = anatolia.DAnatolia安纳托利亚
	f.安纳托利亚Anatolia.SetParent(f)

	f.布凯拉里翁Bucellarian = bucellarian.DBucellarian布凯拉里翁
	f.布凯拉里翁Bucellarian.SetParent(f)

	f.卡帕多西亚Cappadocia = cappadocia.DCappadocia卡帕多西亚
	f.卡帕多西亚Cappadocia.SetParent(f)

	f.哈耳西亚农Charsianon = charsianon.DCharsianon哈耳西亚农
	f.哈耳西亚农Charsianon.SetParent(f)

	f.基比拉奥特Cibyrrhaeot = cibyrrhaeot.DCibyrrhaeot基比拉奥特
	f.基比拉奥特Cibyrrhaeot.SetParent(f)

	f.塞浦路斯Cyprus = cyprus.DCyprus塞浦路斯
	f.塞浦路斯Cyprus.SetParent(f)

	f.奥普西金Nikaea = nikaea.DNikaea奥普西金
	f.奥普西金Nikaea.SetParent(f)

	f.萨摩斯Samos = samos.DSamos萨摩斯
	f.萨摩斯Samos.SetParent(f)

	f.色雷斯西亚Thracesia = thracesia.DThracesia色雷斯西亚
	f.色雷斯西亚Thracesia.SetParent(f)

}
