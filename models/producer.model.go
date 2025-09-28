package models

import "mdx/utils"

type Producer struct {
	ID        *int
	Name      *string `form:"name"`
	Mobile    *string `form:"mobile"`
	Address   *string `form:"address"`
	CreatedOn *string
}

func (producer Producer) Get(offset, limit int, params map[string]string) (CountResult, []any, error) {
	var fields = get_producer_fields(&producer)
	var query string = get_producers_query
	filter_query, filter := producer.QueryBuilder(params)
	query += filter_query
	return GetListData(filter, fields, &producer, "", query, offset, limit)
}

func (producer Producer) GetById(id string) ([]any, error) {
	var fields = get_producer_fields(&producer)
	var query string = get_producers_query + id_filter
	var filter = []any{id}
	return GetSingleData(filter, fields, &producer, query)
}

func (producer *Producer) Add() bool {
	var fields = get_producer_add_fields(producer)
	return utils.RunInsert(insert_producer_query, fields)
}

func (producer *Producer) Update(id string) bool {
	var fields = get_producer_add_fields(producer)
	return utils.RunUpdate(update_producer_query, fields)
}

func (producer Producer) Delete(id string) bool {
	return utils.RunUpdate(delete_producer_query, []any{id})
}

func (producer Producer) QueryBuilder(params map[string]string) (string, []any) {
	var query string
	var filter []any
	if len(params["name"]) > 0 {
		query += name_filter
		filter = append(filter, "%"+params["name"]+"%")
	}
	if len(params["mobile"]) > 0 {
		query += mobile_filter
		filter = append(filter, params["mobile"])
	}
	return query, filter
}
