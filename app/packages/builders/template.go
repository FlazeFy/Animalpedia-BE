package builders

func GetTemplateSelect(name string, firstTable, secondTable *string) string {
	if name == "content_info" {
		return *firstTable + "_slug," + *firstTable + "_name"
	} else if name == "properties_time" {
		return "created_at, updated_at"
	} else if name == "properties_full" {
		return "created_at, created_by, updated_at, updated_by"
	}
	return ""
}

func GetTemplateConcat(name, col string) string {
	if name == "value_group" {
		return "GROUP_CONCAT(" + col + " SEPARATOR ', ') as " + col
	}
	return ""
}

func GetTemplateOrder(name, tableName, ext string) string {
	if name == "permanent_data" {
		return tableName + ".created_at DESC, " + tableName + "." + ext + " DESC "
	} else if name == "dynamic_data" {
		return tableName + ".updated_at DESC, " + tableName + ".created_at DESC, " + tableName + "." + ext + " DESC "
	} else if name == "most_used_normal" {
		return " COUNT(1) DESC"
	}
	return ""
}

func GetTemplateJoin(typeJoin, firstTable, firstCol, secondTable, secondCol string, nullable bool) string {
	var join = ""
	if nullable {
		join = "LEFT JOIN "
	} else {
		join = "JOIN "
	}
	if typeJoin == "same_col" {
		return join + secondTable + " USING(" + firstCol + ") "
	} else if typeJoin == "total" {
		return join + secondTable + " ON " + secondTable + "." + secondCol + " = " + firstTable + "." + firstCol + " "
	}
	return ""
}

func GetTemplateGroup(is_multi_where bool, col string) string {
	var ext = " WHERE "
	if is_multi_where {
		ext = " AND "
	}

	return ext + col + " IS NOT NULL AND trim(" + col + ") != '' GROUP BY " + col + " "
}

func GetTemplateLogic(name string) string {
	if name == "active" {
		return ".deleted_at IS NULL "
	} else if name == "trash" {
		return ".deleted_at IS NOT NULL "
	}
	return ""
}

func GetWhereMine(token string) string {
	return "users_tokens.token ='" + token + "'"
}
