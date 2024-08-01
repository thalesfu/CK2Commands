package feudal

import (
	"bufio"
	"github.com/thalesfu/CK2Commands/earth"
	"github.com/thalesfu/CK2Commands/utils"
	"github.com/thalesfu/paradoxtools/CK2/feud"
	"log"
)

type Title struct {
	Feuds  []feud.Feud
	Holder int
}

func GiveTitle(titles []*Title) {
	file, err := utils.OpenFile("title")

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, title := range titles {
		for _, feud := range title.Feuds {
			writeGiveTitle(writer, title.Holder, feud)
		}
	}
}

func BuildTitle() {
	var titles []*Title
	titles = append(titles, &Title{
		Holder: 2703149,
		Feuds: []feud.Feud{
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CYopurga岳普湖(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CUchturpan乌什吐鲁番(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CUchturpan乌什吐鲁番().BUqturpan乌什(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CKashgar可失合儿().BKashgar可失合儿(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CYopurga岳普湖().BYopurga岳普湖(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CArtux阿图什().BAtush阿图什(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CYarkand鸦儿看(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CKarghalik喀格勒克(),`
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CCadota精绝(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CCherchen车尔臣(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CYarkand鸦儿看().BYarkand鸦儿看(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CKarghalik喀格勒克(),
			//earth.Turkestan图兰.KKhotan于阗().DKarashar喀喇沙尔().CKubera俱毗罗().BKubera俱毗罗(),
			//earth.Turkestan图兰.KKhiva河中().DKhiva希瓦(),
			//earth.Turkestan图兰.KKhiva河中().DKhiva希瓦().CDashhowuz达什霍武兹().BDashhowuz达什霍武兹(),
			//earth.Turkestan图兰.KKhiva河中().DKhuttal珂咄罗().CVakhan镬侃(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳().CKhaylam海拉姆(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳().CFergana费尔干纳(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳().CFergana费尔干纳().BUzkand讹迹邗(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳().CKhaylam海拉姆(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳().CKhojand苦盏(),
			//earth.Turkestan图兰.KKhiva河中().DSamarkand撒马尔罕().CBukhara布哈拉(),
			//earth.Turkestan图兰.KZhetysu七河().DChuy垂河(),
			//earth.Turkestan图兰.KTurkestan乌古斯().DUsturt乌斯秋尔特(),
			//earth.Turkestan图兰.KTurkestan乌古斯().DUsturt乌斯秋尔特().CKusbulak库斯布拉克().BKusbulak库斯布拉克(),
			//earth.Tibet吐蕃.r().DPamir播密().CPamir播密().BKala_panja喀喇喷赤(),
			//earth.Tibet吐蕃.KKashmir迦湿弥罗().DPamir播密().CTashkurgan塔什库尔干(),
			//earth.Persia波斯帝国.KKhorasan呼罗珊().DBalkh缚喝().CBalkh缚喝(),
			//earth.Persia波斯帝国.KKhorasan呼罗珊().DBalkh缚喝().CBalkh缚喝().BTiliatepe蒂拉丘地(),
			//earth.Persia波斯帝国.KKhorasan呼罗珊().DBalkh缚喝().CBalkh缚喝().BAlkhanoum艾哈努姆(),
			//earth.Persia波斯帝国.KKhorasan呼罗珊(),
			//earth.Persia波斯帝国.KDaylam德莱木().DDihistan大益斯坦().CKara_kum法拉瓦().BFarava法拉瓦(),
			//
			//earth.Persia波斯帝国.KDaylam德莱木().DDihistan大益斯坦().CKara_kum法拉瓦(),
			earth.Persia波斯帝国.KAfghanistan迦布罗().DZabulistan社护罗萨他那().CZamindawar扎敏达瓦尔().BKhalai哈莱(),
		},
	})
	//titles = append(titles, &Title{
	//	Holder: 2728929,
	//	Feuds: []feud.Feud{
	//		earth.Persia波斯帝国.KDaylam德莱木().DDihistan大益斯坦().CKara_kum法拉瓦(),
	//	},
	//})
	//titles = append(titles, &Title{
	//	Holder: 2700333,
	//	Feuds: []feud.Feud{
	//		earth.Persia波斯帝国.KAfghanistan迦布罗().DKabul迦布罗().CKabul迦布罗(),
	//	},
	//})
	//titles = append(titles, &Title{
	//	Holder: 2688515,
	//	Feuds: []feud.Feud{
	//		earth.Persia波斯帝国.KAfghanistan迦布罗().DKabul迦布罗().CKunduz昆都士(),
	//	},
	//})
	GiveTitle(titles)
}
