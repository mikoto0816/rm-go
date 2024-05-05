package dao

import (
	"log"
	models "rm-go-blog/modles"
)

func GetPostPage(page, pageSize int) ([]models.Post, error) {

	//分页： page = （page-1）*size
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,? ", page, pageSize)

	var posts []models.Post
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetAllPost() ([]models.Post, error) {

	rows, err := DB.Query("select * from blog_post")
	var posts []models.Post
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func CountGetAllPostCount() (count int) {
	row := DB.QueryRow("select count(1) from blog_post")
	_ = row.Scan(&count)
	return
}
func CountGetAllPostCountBySlug(slug string) (count int) {
	row := DB.QueryRow("select count(1) from blog_post where slug=?", slug)
	_ = row.Scan(&count)
	return
}
func CountGetAllPostByCategoryId(cId int) (count int) {
	row := DB.QueryRow("select count(1) from blog_post where category_id=?", cId)
	_ = row.Scan(&count)
	return
}

func GetPostPageByCategoryId(cid, page, pageSize int) ([]models.Post, error) {

	//分页： page = （page-1）*size
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id=? limit ?,? ", cid, page, pageSize)

	var posts []models.Post
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostDetail(pid int) (models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid = ?", pid)
	var post models.Post
	if row.Err() != nil {
		return post, row.Err()
	}
	err := row.Scan(&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		log.Println(err)
		return post, err
	}
	return post, nil
}

func SavePost(post *models.Post) {

	insert, err := DB.Exec("insert into blog_post (title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values (?,?,?,?,?,?,?,?,?,?);", post.Title, post.Content, post.Markdown, post.CategoryId, post.UserId, post.ViewCount,
		post.Type, post.Slug, post.CreateAt, post.UpdateAt)
	if err != nil {
		log.Println(err)
	}
	pid, _ := insert.LastInsertId()
	post.Pid = int(pid)
}

func UpdatePost(post *models.Post) {

	_, err := DB.Exec("update blog_post set title =?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid = ?",
		post.Title, post.Content, post.Markdown, post.CategoryId,
		post.Type, post.Slug, post.UpdateAt, post.Pid)
	if err != nil {
		log.Println(err)
	}

}

func GetPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {

	//分页： page = （page-1）*size
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where slug=? limit ?,? ", slug, page, pageSize)

	var posts []models.Post
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func Search(condtion string) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post where title like ?", "%"+condtion+"%")
	var posts []models.Post
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
