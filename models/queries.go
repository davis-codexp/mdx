package models

var id_filter = " AND id = ?"
var name_filter = " AND name LIKE ?"
var mobile_filter = " AND mobile = ?"
var genre_filter = " AND genre = ?"
var order_filter = " ORDER BY created_on DESC"

/*-------------PRODUCER QUERIES---------*/
var get_producers_query = `
	SELECT id, name, mobile, address, created_on
	FROM producers WHERE is_deleted = 0
`
var count_producers_query = `
	SELECT COUNT(id) FROM producers WHERE is_deleted = 0
`
var insert_producer_query = `
	INSERT INTO producers(name, mobile, address) VALUES(?, ?, ?)
`
var update_producer_query = `
	UPDATE producers SET name = ?, mobile = ?, address = ?
	WHERE id = ?
`
var delete_producer_query = `
	UPDATE producers SET is_deleted = 1 WHERE id = ?
`

/*-------------MOVIE QUERIES---------*/
var get_movies_query = `
	SELECT movies.id, movies.name, genre, length, year, synopsis,
	price, manifest, producer, producers.name AS producer_name,
	thumbnail, movies.created_on
	FROM movies
	INNER JOIN producers ON producers.id = movies.producer
	WHERE movies.is_deleted = 0
`
var count_movies_query = `
	SELECT COUNT(movies.id) FROM movies WHERE is_deleted = 0
`
var insert_movie_query = `
	INSERT INTO movies(
		name, genre, length, year, synopsis, price,
		manifest, producer, thumbnail
	) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)
`
var update_movie_query = `
	UPDATE movies SET name = ?, genre = ?, length = ?,
	year = ?, synopsis = ?, price = ?,
	manifest = ?, producer = ?, thumbnail = ?
	WHERE id = ?
`
var delete_movie_query = `
	UPDATE movies SET is_deleted = 1 WHERE id = ?
`
var movies_id_filter = " AND movies.id = ?"
var movies_name_fitler = " AND movies.name LIKE ?"
var movies_order_filter = " ORDER BY movies.created_on DESC"

/*-------------ACTOR QUERIES---------*/
var get_actors_query = `
	SELECT id, name, image
	FROM actors WHERE is_deleted = 0
`
var count_actors_query = `
	SELECT COUNT(id) FROM actors WHERE is_deleted = 0
`
var insert_actor_query = `
	INSERT INTO actors(name, image) VALUES(?, ?)
`
var update_actor_query = `
	UPDATE actors SET name = ?, image = ?
	WHERE id = ?
`
var delete_actor_query = `
	UPDATE actors SET is_deleted = 1 WHERE id = ?
`

/*-------------CUSTOMER QUERIES---------*/
var get_customers_query = `
	SELECT id, name, mobile
	FROM customers WHERE is_deleted = 0
`
var count_customers_query = `
	SELECT COUNT(id) FROM customers WHERE is_deleted = 0
`
var insert_customer_query = `
	INSERT INTO customers(name, mobile) VALUES(?, ?)
`
var update_customer_query = `
	UPDATE customers SET name = ?, mobile = ?
	WHERE id = ?
`
var delete_customer_query = `
	UPDATE customers SET is_deleted = 1 WHERE id = ?
`

/*-------------MOVIE CASTS QUERIES---------*/
var get_movie_casts_query = `
	SELECT id, name, mobile
	FROM customers WHERE is_deleted = 0
`
var insert_movie_casts_query = `
	INSERT INTO movie_casts(movie, actor) VALUES
`
var on_duplicate_movie_cast = `
	ON DUPLICATE KEY UPDATE movie = VALUES(movie),
	actor = VALUES(actor)
`
var delete_movie_cast_query = "DELETE FROM movie_casts WHERE movie = ? AND actor = ?"
