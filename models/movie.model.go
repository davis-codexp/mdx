package models

import "mdx/utils"

type Movie struct {
	ID           *int         `json:"id"`
	Name         *string      `form:"name" json:"name"`
	Genre        *string      `form:"genre" json:"genre"`
	Length       *int         `form:"length" json:"length"`
	Year         *int         `form:"year" json:"year"`
	Synopsis     *string      `form:"synopsis" json:"synopsis"`
	Price        *int         `form:"price" json:"price"`
	Producer     *int         `form:"producer" json:"producer"`
	Thumbnail    *string      `form:"thumbnail" json:"thumbnail"`
	CreatedOn    *string      `json:"created_on"`
	ProducerName *string      `json:"producer_name"`
	Cast         *[]MovieCast `form:"cast" json:"cast"`
}

type MovieCast struct {
	Movie     *int    `form:"movie" json:"movie"`
	Actor     *int    `form:"actor" json:"actor"`
	ActorName *string `json:"actor_name"`
	Image     *string `json:"image"`
}

type Actor struct {
	ID        *int
	Name      *string `form:"name" json:"name"`
	Image     *string `form:"image" json:"image"`
	CreatedOn *string `json:"created_on"`
}

/*----------------------------MOVIE-------------------------*/
func (obj Movie) Get(offset, limit int, params map[string]string) (CountResult, []any, error) {
	var fields = get_movie_fields(&obj)
	var query string = get_movies_query
	filter_query, filter := obj.QueryBuilder(params)
	query += filter_query + movies_order_filter
	return GetListData(filter, fields, &obj, "", query, offset, limit)
}

func (obj Movie) GetById(id string) ([]any, error) {
	var fields = get_movie_fields(&obj)
	var query string = get_movies_query + movies_id_filter
	var filter = []any{id}
	return GetSingleData(filter, fields, &obj, query)
}

func (obj *Movie) Add() bool {
	var fields = get_movie_add_fields(obj)
	return utils.RunInsert(insert_movie_query, fields)
}

func (obj *Movie) Update(id string) bool {
	var fields = get_movie_add_fields(obj)
	fields = append(fields, id)
	return utils.RunUpdate(update_movie_query, fields)
}

func (obj Movie) Delete(id string) bool {
	return utils.RunUpdate(delete_movie_query, []any{id})
}

func (obj Movie) QueryBuilder(params map[string]string) (string, []any) {
	var query string
	var filter []any
	if len(params["name"]) > 0 {
		query += movies_name_fitler
		filter = append(filter, "%"+params["name"]+"%")
	}
	if len(params["genre"]) > 0 {
		query += genre_filter
		filter = append(filter, params["genre"])
	}
	return query, filter
}

/*--------------------------------MOVIE-------------------------*/
func (obj Actor) Get(offset, limit int, params map[string]string) (CountResult, []any, error) {
	var fields = get_actor_fields(&obj)
	var query string = get_actors_query
	filter_query, filter := obj.QueryBuilder(params)
	query += filter_query + order_filter
	return GetListData(filter, fields, &obj, "", query, offset, limit)
}

func (obj Actor) GetById(id string) ([]any, error) {
	var fields = get_actor_fields(&obj)
	var query string = get_actors_query + id_filter
	var filter = []any{id}
	return GetSingleData(filter, fields, &obj, query)
}

func (obj *Actor) Add() bool {
	var fields = get_actor_add_fields(obj)
	return utils.RunInsert(insert_actor_query, fields)
}

func (obj *Actor) Update(id string) bool {
	var fields = get_actor_add_fields(obj)
	fields = append(fields, id)
	return utils.RunUpdate(update_actor_query, fields)
}

func (obj Actor) Delete(id string) bool {
	return utils.RunUpdate(delete_actor_query, []any{id})
}

func (obj Actor) QueryBuilder(params map[string]string) (string, []any) {
	var query string
	var filter []any
	if len(params["name"]) > 0 {
		query += movies_name_fitler
		filter = append(filter, "%"+params["name"]+"%")
	}
	return query, filter
}
