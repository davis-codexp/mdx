package models

func get_producer_fields(obj *Producer) []any {
	return []any{
		&obj.ID, &obj.Name, &obj.Mobile, &obj.Address, &obj.CreatedOn,
	}
}

func get_producer_add_fields(obj *Producer) []any {
	return []any{obj.Name, obj.Mobile, obj.Address}
}

func get_movie_fields(obj *Movie) []any {
	return []any{
		&obj.ID, &obj.Name, &obj.Genre, &obj.Length,
		&obj.Year, &obj.Synopsis, &obj.Price,
		&obj.Producer, &obj.ProducerName, &obj.Thumbnail,
		&obj.CreatedOn,
	}
}

func get_movie_add_fields(obj *Movie) []any {
	return []any{
		obj.Name, obj.Genre, obj.Length, obj.Year, obj.Synopsis,
		obj.Price, obj.Producer, obj.Thumbnail,
	}
}

func get_actor_fields(obj *Actor) []any {
	return []any{&obj.ID, &obj.Name, &obj.Image}
}

func get_actor_add_fields(obj *Actor) []any {
	return []any{obj.Name, obj.Image}
}
