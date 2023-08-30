package aggregatepipeline

import (
	"client_administration/constants"
)

func GetUserDataForAdminPipeline (searchFilter constants.SearchUserData) []map[string]interface{}{
	// id := ""
	// if searchFilter.Id != ""{
	// 	id = searchFilter.Id
	// }
	search := ""
	if searchFilter.Search != "" {
		search = searchFilter.Search
	}
	active := searchFilter.Active

	// IdObject, _ := utils.CreatObjectID(id)

    pipeline := []map[string]interface{}{}

	pipeline = append(pipeline, map[string]interface{}{
		"$match": map[string]interface{}{
			"active": active,
		},
	})

    pipeline = append(pipeline, map[string]interface{}{
        "$project": map[string]interface{}{
            "email":           1,
            "firstname":       1,
            "lastname":        1,
            "active":          1,
            "_id":         		1,
            "telephone":   		1,
            "role":            	1,
            "postaladdress":    1,
        },
    })

    pipeline = append(pipeline, map[string]interface{}{
        "$match": map[string]interface{}{
            "$or": []map[string]interface{}{
                {"postaladdress.postalcode":     map[string]interface{}{"$regex": search, "$options": "i"}},
                {"postaladdress.city":      map[string]interface{}{"$regex": search, "$options": "i"}},
                {"postaladdress.country":      map[string]interface{}{"$regex": search, "$options": "i"}},
                {"lastname":  map[string]interface{}{"$regex": search, "$options": "i"}},
            },
        },
    })

    return pipeline
}
