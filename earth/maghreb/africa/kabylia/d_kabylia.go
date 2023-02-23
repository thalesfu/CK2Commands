package kabylia

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/kabylia/annaba"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/kabylia/bejaija"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/kabylia/biskra"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/kabylia/constantine"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/kabylia/tebessa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KabyliaDuke interface {
	feud.Duke
	CAnnaba安纳巴() annaba.AnnabaCounty
	CBejaija贝贾亚() bejaija.BejaijaCounty
	CBiskra比斯卡拉() biskra.BiskraCounty
	CConstantine君士坦丁() constantine.ConstantineCounty
	CTebessa泰贝萨() tebessa.TebessaCounty
}

type 卡比利亚KabyliaDuke struct {
	feud.BaseDuke
	安纳巴Annaba       annaba.AnnabaCounty
	贝贾亚Bejaija      bejaija.BejaijaCounty
	比斯卡拉Biskra      biskra.BiskraCounty
	君士坦丁Constantine constantine.ConstantineCounty
	泰贝萨Tebessa      tebessa.TebessaCounty
}

func (d *卡比利亚KabyliaDuke) CAnnaba安纳巴() annaba.AnnabaCounty {
	return d.安纳巴Annaba
}

func (d *卡比利亚KabyliaDuke) CBejaija贝贾亚() bejaija.BejaijaCounty {
	return d.贝贾亚Bejaija
}

func (d *卡比利亚KabyliaDuke) CBiskra比斯卡拉() biskra.BiskraCounty {
	return d.比斯卡拉Biskra
}

func (d *卡比利亚KabyliaDuke) CConstantine君士坦丁() constantine.ConstantineCounty {
	return d.君士坦丁Constantine
}

func (d *卡比利亚KabyliaDuke) CTebessa泰贝萨() tebessa.TebessaCounty {
	return d.泰贝萨Tebessa
}

var DKabylia卡比利亚 KabyliaDuke = &卡比利亚KabyliaDuke{}

func init() {
	f := DKabylia卡比利亚.(*卡比利亚KabyliaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kabylia",
		TitleName: "卡比利亚",
		TitleCode: "d_kabylia",
		Counties:  map[string]feud.County{},
	}

	f.安纳巴Annaba = annaba.CAnnaba安纳巴
	f.安纳巴Annaba.SetParent(f)

	f.贝贾亚Bejaija = bejaija.CBejaija贝贾亚
	f.贝贾亚Bejaija.SetParent(f)

	f.比斯卡拉Biskra = biskra.CBiskra比斯卡拉
	f.比斯卡拉Biskra.SetParent(f)

	f.君士坦丁Constantine = constantine.CConstantine君士坦丁
	f.君士坦丁Constantine.SetParent(f)

	f.泰贝萨Tebessa = tebessa.CTebessa泰贝萨
	f.泰贝萨Tebessa.SetParent(f)

}
