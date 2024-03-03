// pkg/models/mysql/news.go

package mysql

import (
	_ "github.com/go-sql-driver/mysql"
)

// News представляет структуру новости.
type News struct {
	ID       int
	Title    string
	Content  string
	Tag      string
	ImageURL string
}

// CreateNews создает новость в базе данных.
func (dbm *DBModel) CreateNews(title, content, tag, imageURL string) (int, error) {
	result, err := dbm.DB.Exec("INSERT INTO news (title, content, tag, image_url) VALUES (?, ?, ?, ?)", title, content, tag, imageURL)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// GetNews возвращает список всех новостей из базы данных.
func (dbm *DBModel) GetNews() ([]News, error) {
	rows, err := dbm.DB.Query("SELECT id, title, content, tag, image_url FROM news")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList []News
	for rows.Next() {
		var news News
		err := rows.Scan(&news.ID, &news.Title, &news.Content, &news.Tag, &news.ImageURL)
		if err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}

	return newsList, nil
}
func (dbm *DBModel) GetNewsByCategory(s string) ([]News, error) {
	rows, err := dbm.DB.Query("SELECT id, title, content, tag, image_url FROM news where tag = ?", s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList []News
	for rows.Next() {
		var news News
		err := rows.Scan(&news.ID, &news.Title, &news.Content, &news.Tag, &news.ImageURL)
		if err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}

	return newsList, nil
}

// DeleteNews удаляет новость по ее идентификатору из базы данных.
func (dbm *DBModel) DeleteNews(title string) error {
	_, err := dbm.DB.Exec("DELETE FROM news WHERE title = ?", title)
	return err
}

func (dbm *DBModel) UpdateNews(oldtitle, title, content, tag, imageURL string) error {
	_, err := dbm.DB.Exec("UPDATE news SET title = ?, content = ?, tag = ? where title = ?", title, content, tag, oldtitle)
	return err
}

// NewsExists checks if a news item with the specified title exists in the database.
func (dbm *DBModel) NewsExists(title string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM news WHERE title = ? LIMIT 1)"
	err := dbm.DB.QueryRow(query, title).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

// IsTitleDuplicate checks if a title already exists in the database, excluding the news with the specified oldTitle.
func (dbm *DBModel) IsTitleDuplicate(newTitle, oldTitle string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM news WHERE title = ? AND title != ? LIMIT 1)"
	err := dbm.DB.QueryRow(query, newTitle, oldTitle).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
