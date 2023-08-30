package aggregatepipeline

import (
	"client_administration/constants"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUserDataForAdminPipeline (searchFilter constants.SearchUserData) []map[string]interface{}{
	
    // Search if search params is not empty
	search := ""
	if searchFilter.Search != "" {
		search = searchFilter.Search
	}
	active := searchFilter.Active
    

    pipeline := []map[string]interface{}{}

    // Get all document where user is active
	pipeline = append(pipeline, map[string]interface{}{
		"$match": map[string]interface{}{
			"active": active,
		},
	})

    // Return the listed fields only
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

    // Search key that will return value by
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
    
    // Pagenation by skip and limit
    skipNumber := searchFilter.ExtraParams.Skip
    limitNumber := searchFilter.ExtraParams.Limit
        if skipNumber != 0 {
            skip := bson.M{"$skip": skipNumber}
            pipeline = append(pipeline, skip)
        }
        if limitNumber != 0 {
            limit := bson.M{"$limit": limitNumber}
            pipeline = append(pipeline, limit)
        }

    return pipeline
}
