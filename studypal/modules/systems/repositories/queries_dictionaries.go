package repositories

import (
	"database/sql"
	"net/http"
	"studypal/modules/systems/models"
	"studypal/packages/builders"
	"studypal/packages/database"
	"studypal/packages/helpers/converter"
	"studypal/packages/helpers/generator"
	"studypal/packages/helpers/response"
)

func GetDictionary(path string, types *string, is_active bool) (response.Response, error) {
	// Declaration
	var obj models.GetAllDictionary
	var arrobj []models.GetAllDictionary
	var res response.Response
	var baseTable = "dictionary"
	var secondTable = "dictionary_type"
	var sqlStatement string
	var whereType string

	// Nullable column
	var DctDesc sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)

	if is_active {
		whereType = "WHERE " + secondTable + ".dctt_name = '" + *types + "' AND is_active = 1 "
	} else {
		whereType = " "
	}

	order := builders.GetTemplateOrder("dynamic_data", baseTable, "dictionary_name")
	join1 := builders.GetTemplateJoin("total", baseTable, "dictionary_type", secondTable, "id", false)

	sqlStatement = "SELECT " + selectTemplate + " " +
		"FROM " + baseTable + " " +
		join1 + " " + whereType +
		"ORDER BY " + order

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.DctSlug,
			&obj.DctName,
			&DctDesc,
		)

		if err != nil {
			return res, err
		}

		obj.DctDesc = converter.CheckNullString(DctDesc)

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, &whereType)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("dictionary", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
