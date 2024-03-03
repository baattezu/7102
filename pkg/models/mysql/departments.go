package mysql

////////////////////////////////////////// DEPARTMENTS /////////////////////////////////////////

type Department struct {
	ID        int
	Dep_name  string
	Staff_qnt int
}

func (dbm *DBModel) CreateDepartment(dep_name string, staff_quantity int) (int, error) {
	result, err := dbm.DB.Exec("INSERT INTO departments (dep_name, staff_quantity) VALUES (?, ?)", dep_name, staff_quantity)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (dbm *DBModel) GetDepartments() ([]Department, error) {
	rows, err := dbm.DB.Query("SELECT id, dep_name, staff_quantity FROM departments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departmentList []Department
	for rows.Next() {
		var departments Department
		err := rows.Scan(&departments.ID, &departments.Dep_name, &departments.Staff_qnt)
		if err != nil {
			return nil, err
		}
		departmentList = append(departmentList, departments)
	}

	return departmentList, nil
}
func (dbm *DBModel) GetDepByCategory(s string) ([]News, error) {
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

func (dbm *DBModel) DeleteDepartment(id int) error {
	_, err := dbm.DB.Exec("DELETE FROM department WHERE title = ?", id)
	return err
}

func (dbm *DBModel) UpdateDepartment(id int, newname string, newstaff int) error {
	_, err := dbm.DB.Exec("UPDATE department SET dep_name = ?, staff_quantity = ? where id = ?", newname, newstaff, id)
	return err
}
