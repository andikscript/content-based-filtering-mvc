package service

import (
	"contentbasedfilteringmvc/cmd/database"
	"contentbasedfilteringmvc/cmd/model"
	"contentbasedfilteringmvc/cmd/util"
	"fmt"
	"strings"
	"time"
)

func ProcessAlgorithm(text string, idUser string) []model.QueryProduct {
	saveHistory := model.SaveHistory{
		Created:     time.Now().Format("2006-01-02 15:04:05"),
		Updated:     time.Now().Format("2006-01-02 15:04:05"),
		IdHistory:   util.RandomNumber(20),
		IdUser:      idUser,
		Description: strings.ToLower(text),
	}

	splt := strings.Split(strings.ToLower(strings.TrimSpace(text)), " ")

	isValid, resultProduct, merkProduct, hargaProduct := getAllProduct()
	if !isValid {
		fmt.Println("⛔ no data found")
	}

	filter := filteringProduct(resultProduct, splt)

	saveCondition := false

	if len(filter) == 0 && !strings.Contains(text, "jt") && !strings.Contains(text, "murah") {
		fmt.Println("⛔ no data found")

		isValid, desc := database.ReadHistory(idUser)
		fmt.Println("desc", desc)
		if isValid {
			fmt.Println("➡️ rekomendasi pencarian terakhir..")
			text = desc
			splt = strings.Split(text, " ")
			saveCondition = true
		}
	}

	// found harga
	if isValid := strings.Contains(text, "jt") || strings.Contains(text, "murah"); isValid {
		price := getPrice(splt)
		group := []string{}

		for k, v := range hargaProduct {
			if v >= price && v <= price+1000000 {
				group = append(group, k)
			}
		}

		if !saveCondition {
			// save history
			database.InsertHistory(saveHistory)
		}

		// print
		return showProduct(group)
	}

	// found merk
	for _, merk := range merkProduct {
		for i := 0; i < len(splt); i++ {
			if strings.Contains(merk, splt[i]) {
				isAvail, resultByMerk := getProductByMerk(merk)

				if !isAvail {
					fmt.Println("no data")
					return []model.QueryProduct{}
				}

				// filter
				filter := filteringProduct(resultByMerk, splt)

				// sort
				sorting := sortResult(filter)

				if !saveCondition {
					// save history
					database.InsertHistory(saveHistory)
				}

				// print
				return showProduct(sorting)
			}
		}
	}

	// sort
	sorting := sortResult(filter)

	if !saveCondition {
		// save history
		database.InsertHistory(saveHistory)
	}

	// print
	return showProduct(sorting)
}
