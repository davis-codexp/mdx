package models

func get_producer_fields(obj *Producer) []any {
	return []any{
		&obj.ID, &obj.Name, &obj.Mobile, &obj.Address, &obj.CreatedOn,
	}
}

func get_producer_add_fields(obj *Producer) []any {
	return []any{
		obj.Name, obj.Mobile, obj.Address,
	}
}
