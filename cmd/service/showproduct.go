package service

import (
	"contentbasedfilteringmvc/cmd/database"
	"contentbasedfilteringmvc/cmd/model"
	"fmt"
	"sort"
)

func showProduct(maps []string) []model.QueryProduct {
	if len(maps) == 0 {
		return []model.QueryProduct{}
	}

	rs := []model.QueryProduct{}

	for _, v := range maps {
		isValid, product := database.ReadProductById(v)

		if isValid {
			//pr := fmt.Sprintf("📲 %s 📍 Tipe %s 📀 RAM %sGB 💾 Internal %sGB 📸 Kamera Utama %sMP dibandrol dengan harga 💸Rp.%d",
			//	product.Merk, product.Type, product.Ram, product.Internal, product.Camera, product.Harga)
			//fmt.Println(pr)

			rs = append(rs, model.QueryProduct{
				Merk:     product.Merk,
				Type:     product.Type,
				Ram:      fmt.Sprintf("%sgb", product.Ram),
				Internal: fmt.Sprintf("%sgb", product.Internal),
				Camera:   fmt.Sprintf("%sMP", product.Camera),
				Harga:    product.Harga,
			})
		}
	}

	return rs
}

func sortResult(maps map[string]int) []string {
	var ss []kv
	for k, v := range maps {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	keys := []string{}
	for _, kv := range ss {
		keys = append(keys, kv.Key)
	}

	return keys
}

type kv struct {
	Key   string
	Value int
}
